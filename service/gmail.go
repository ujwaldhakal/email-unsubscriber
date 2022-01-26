package service

import (
	"fmt"
	"google.golang.org/api/gmail/v1"
	"log"
	"os"
	"github.com/joho/godotenv"
	"time"
)

type PageToken struct {
	Token string
}

func GetMessageList(srv *gmail.Service, userId string, token PageToken) *gmail.ListMessagesResponse {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dateFrom := fmt.Sprintf("%d" ,convertDateToTimestamp(os.Getenv("SEARCH_DATE_FROM")))
	searchQuery  := os.Getenv("INBOX_SEARCH_QUERY")

	fmt.Println("got it search",searchQuery)
	fmt.Println("got it date",dateFrom)
	data, err := srv.Users.Messages.List(userId).Q(searchQuery+ " after:" + dateFrom).PageToken(token.Token).Do()
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
