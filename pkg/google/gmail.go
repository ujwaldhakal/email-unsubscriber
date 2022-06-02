package google

import (
	"fmt"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
	"log"
	"time"
)

type PageToken struct {
	Token string
}

	type GmailApi interface {
		Q(q string) *gmail.UsersMessagesListCall
		PageToken(pageToken string) *gmail.UsersMessagesListCall
		Do(opts ...googleapi.CallOption) (*gmail.ListMessagesResponse, error)
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
	fmt.Println(srv.PageToken(gmail.Token))
	fmt.Println("yolo")
	data, err := srv.Q(gmail.SearchQuery+ " after:" + dateFrom).PageToken(gmail.Token).Do()
	fmt.Println("not here la")
	if err != nil {
		fmt.Println("got here la")
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
