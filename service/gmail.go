package service

import (
	"google.golang.org/api/gmail/v1"
	"log"
	"os"
)

type PageToken struct {
	Token string;
}

func GetMessageList(srv *gmail.Service,  userId string, token PageToken) *gmail.ListMessagesResponse  {

	dateFrom := os.Getenv("SEARCH_DATE_FROM")
	data, err := srv.Users.Messages.List(userId).Q("label=promotions after:"+dateFrom).PageToken(token.Token).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}

	return data

}