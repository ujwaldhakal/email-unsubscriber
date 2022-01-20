package cmd

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/ujwaldhakal/email-unsubscriber/pkg/pgsql"
	service "github.com/ujwaldhakal/email-unsubscriber/model"
	"github.com/ujwaldhakal/email-unsubscriber/pkg/rabbitmq"
	gmailApi "github.com/ujwaldhakal/email-unsubscriber/service"
	"google.golang.org/api/gmail/v1"
	"mvdan.cc/xurls/v2"
	"regexp"
	"strings"
)



var parseEmail = &cobra.Command{
	Use:   "parse-email",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		ctx := context.Background()


		pgsql.GetConnection()

		msgs := rabbitmq.ConsumerClient("parse-email")

		forever := make(chan bool)

		go func() {
			for d := range msgs {

				parseThread(string(d.Body),ctx)
				fmt.Println("data consumed",string(d.Body))

				d.Ack(true)
			}
		}()


		fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever

	},
}



func parseThread(threadId string, ctx context.Context)  {
	srv := gmailApi.GetService(ctx)
	data, e := srv.Users.Messages.Get("me",threadId).Do()
	if e != nil {
		fmt.Println(e)
	}

	unsubscribeLink,SenderLink := findUnsubscribeAndSenderLinkFromHeader(data)


	if unsubscribeLink == "" && len(data.Payload.Parts) > 0 {

		body := data.Payload.Parts[0]

		if len(data.Payload.Parts) > 1 {
			body = data.Payload.Parts[1]
		}

		unsubscribeLink = findUnsubscribeLinkInHtml(body.Body.Data)
	}


	if unsubscribeLink != "" {
		fmt.Println("got it",pgsql.SearchByNameAndSender(parseUrlFromEmail(SenderLink),SenderLink))

		service := service.Service{
			ID: uuid.NewString(),
			Name: parseUrlFromEmail(SenderLink),
			Sender: SenderLink,
			ThreadId: threadId,
			UnsubscribeLink: unsubscribeLink,
			Unsubscribed:    false,
		}

		serviceName := parseUrlFromEmail(SenderLink)

		if len(service.SearchByNameAndSender(serviceName,SenderLink)) > 0 {
			return
		}

		service.Create(service)
	}

}

func findUnsubscribeAndSenderLinkFromHeader(data *gmail.Message) (string,string)  {

	 unsubscribeLink := ""
	 senderEmail := ""
	for _,header := range data.Payload.Headers{
		if header.Name == "List-Unsubscribe"{
			xurlsStrict := xurls.Strict()
			output := xurlsStrict.FindAllString(header.Value, -1)


			unsubscribeLink = output[0]
			if !strings.Contains(unsubscribeLink,"https://") && len(output) > 1 {
				unsubscribeLink = output[1]
			}

			if !strings.Contains(unsubscribeLink,"https://")  {
				unsubscribeLink = ""
			}
		}


		if header.Name == "From"{
			senderEmail = parseEmailFromString(header.Value)
		}
	}
	return unsubscribeLink,senderEmail
}

func parseUrlFromEmail(email string) string {

	splitedStr := strings.Split(email,"@")

	return splitedStr[1]
}

func parseEmailFromString(data string) (string) {
	re := regexp.MustCompile(`<(.*?)\>`)
	matched := re.FindString(data)
	email := strings.TrimLeft(strings.TrimRight(matched, ">"), "<")

	return email
}

func findUnsubscribeLinkInHtml(data string) string {
	uDec, _ := b64.URLEncoding.DecodeString(data)

	html := string(uDec)

	re := regexp.MustCompile("<a\\s+(?:[^>]*?\\s+)?href=\"([^\"]*)\">\\s*(?i)(unsubscribe)")
	matched := re.FindString(html)

	xurlsStrict := xurls.Strict()
	output := xurlsStrict.FindAllString(matched, -1)

	if len(output) == 0 {
		return "";
	}
	return output[0]
}

func init()  {
	rootCmd.AddCommand(parseEmail)
}
