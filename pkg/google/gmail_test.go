package google

import (
	"context"
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/gmail/v1"
	"testing"
)

func TestConvertDateToTimestamp(t *testing.T)  {

	assert.Equal(t,convertDateToTimestamp("2021-01-01"),int64(1609459200));
}

type gmailApi struct {
	expectedErr  error
}


func (mrc *gmailApi) List(userId string) *gmail.UsersMessagesListCall {
	ctx := context.Background()

	srv := GetService(ctx)

	c := &gmail.UsersMessagesListCall{s: srv}
	return c
}

func TestGetMessageList(t *testing.T)  {

	mrc := &gmailApi{}
	gmail := &Gmail{
		Token: "asd",
		UserId: "me",
		SearchDate: "2021-01-01",
		SearchQuery: "q",
	}
	gmail.GetMessageList(mrc)
	assert.Equal(t,convertDateToTimestamp("2021-01-01"),int64(1609459200));
}