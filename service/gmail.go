package service

import (
	"google.golang.org/api/gmail/v1"
	"log"
)

type PageToken struct {
	Token string;
}

func GetMessageList(srv *gmail.Service,  userId string, token PageToken) *gmail.ListMessagesResponse  {

	data, err := srv.Users.Messages.List(userId).Q("unsubscribe after:1642258829").PageToken(token.Token).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}

	return data

}