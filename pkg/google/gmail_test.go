package google

import (
	"errors"
	_ "fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ujwaldhakal/email-unsubscriber/mocks/mock"
	apigmail "google.golang.org/api/gmail/v1"
	"testing"
)

func TestConvertDateToTimestamp(t *testing.T)  {

	assert.Equal(t,convertDateToTimestamp("2021-01-01"),int64(1609459200));
}



func TestGetMessageList(t *testing.T)  {

	gmail := &Gmail{
		Token: "asd",
		UserId: "me",
		SearchDate: "2021-01-01",
		SearchQuery: "q",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockGmailApi(mockCtrl)
	mockDoer.EXPECT().Do(gomock.Any()).Return(nil,errors.New("something is wrong"))
	mockDoer.EXPECT().Q(gomock.Any()).Return(&apigmail.UsersMessagesListCall{})
	mockDoer.EXPECT().PageToken(gomock.Any()).Return(&apigmail.UsersMessagesListCall{})
	//mockDoer.EXPECT().PageToken(gomock.Any()).Return(&gmail.U)
	gmail.GetMessageList(mockDoer)
}