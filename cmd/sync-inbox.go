package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	rabbitmq "github.com/ujwaldhakal/email-unsubscriber/pkg/rabbitmq"
	"github.com/ujwaldhakal/email-unsubscriber/service"
	gmailApi "github.com/ujwaldhakal/email-unsubscriber/service"
	"google.golang.org/api/gmail/v1"
)

type messageId struct {
	id string
}

func getMessages(srv *gmail.Service, token service.PageToken) {

	messages := gmailApi.GetMessageList(srv, "me", token)

	fmt.Println("total message", len(messages.Messages))
	for _, d := range messages.Messages {
		id := d.Id
		//time.Sleep(3 * time.Second)
		rabbitmq.Publish("parse-email", []byte(id))

	}

	if messages.NextPageToken != "" {
		getMessages(srv, service.PageToken{messages.NextPageToken})
	}
}

var syncInbox = &cobra.Command{
	Use: "sync-inbox",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		ctx := context.Background()

		srv := gmailApi.GetService(ctx)

		getMessages(srv, service.PageToken{})

	},
}

func init() {
	rootCmd.AddCommand(syncInbox)
}
