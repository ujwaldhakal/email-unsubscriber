package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	mysql "github.com/ujwaldhakal/email-unsubscriber/pkg/pgsql"
	rabbitmq "github.com/ujwaldhakal/email-unsubscriber/pkg/rabbitmq"
	"github.com/ujwaldhakal/email-unsubscriber/service"
	gmailApi "github.com/ujwaldhakal/email-unsubscriber/service"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
)


type messageId struct {
	id string;
}


func getMessages(srv *gmail.Service, token service.PageToken) {

	messages := gmailApi.GetMessageList(srv,"me",token)


	fmt.Println("total message",len(messages.Messages))
	for _,d := range messages.Messages {
		//m, _ := srv.Users.Messages.Get(user,d.Id).Do()

		//for _,da := range m.Payload.Headers{
		//	fmt.Println(da)
		//}
		//fmt.Println(m.Payload.Parts[0].Headers[0].Value)
		//uDec, _ := b64.URLEncoding.DecodeString(m.Payload.Parts[0].Body.Data)
		//fmt.Println(string(uDec))
		encodedData, err := json.Marshal(d.Id)
		if err != nil {
			log.Fatal("error encoding to json")
		}
		rabbitmq.Publish("test",encodedData)
	}

	if messages.NextPageToken != "" {
		getMessages(srv,service.PageToken{messages.NextPageToken})
	}
}

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		ctx := context.Background()
		b, err := ioutil.ReadFile("credentials.json")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// If modifying these scopes, delete your previously saved token.json.
		config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}


		client := gmailApi.GetClient(config)

		srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Unable to retrieve Gmail client: %v", err)
		}

		fmt.Println("hello",srv)

		mysql.GetConnection()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}