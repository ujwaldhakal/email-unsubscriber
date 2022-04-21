package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ujwaldhakal/email-unsubscriber/pkg/google"
	"github.com/ujwaldhakal/email-unsubscriber/pkg/rabbitmq"
	"google.golang.org/api/gmail/v1"
)

type messageId struct {
	id string
}

func getMessages(srv *gmail.Service, token google.PageToken) {

	fmt.Println("damn",*srv.Users.Messages)
	gmail := &google.Gmail{
		UserId: "me",
		Token: token.Token,
		SearchQuery: "q",
		SearchDate: "2021-01-01",

	}
	messages := gmail.GetMessageList(srv.Users.Messages)
	fmt.Println("total message", len(messages.Messages))
	for _, d := range messages.Messages {
		id := d.Id
		//time.Sleep(3 * time.Second)
		rabbitmq.Publish("parse-email", []byte(id))

	}

	if messages.NextPageToken != "" {
		getMessages(srv, google.PageToken{messages.NextPageToken})
	}
}

var syncInbox = &cobra.Command{
	Use: "sync-inbox",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here

		ctx := context.Background()

		srv := google.GetService(ctx)

		getMessages(srv, google.PageToken{})

	},
}

func init() {
	rootCmd.AddCommand(syncInbox)
}
