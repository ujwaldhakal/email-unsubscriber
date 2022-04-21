package google

import (
	"fmt"
	"google.golang.org/api/gmail/v1"
	"log"
	"time"
)

type PageToken struct {
	Token string
}

type GmailApi interface {
	List(userId string) *gmail.UsersMessagesListCall
}

type Gmail struct {
	UserId string
	Token string
	SearchDate string
	SearchQuery string
}

func (gmail Gmail) GetMessageList(srv GmailApi) *gmail.ListMessagesResponse {
	dateFrom := fmt.Sprintf("%d" , convertDateToTimestamp(gmail.SearchDate))

	fmt.Println("upto here")
	data, err := srv.List(gmail.UserId).Q(gmail.SearchQuery+ " after:" + dateFrom).PageToken(gmail.Token).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve messages: %v", err)
	}

	return data

}

func convertDateToTimestamp(date string) int64 {

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal(err)
	}

	return t.Unix()
}
