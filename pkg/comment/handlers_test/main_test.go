package handlers_test

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/comment"
	"os"
	"testing"
)

func NewMockResponder(rc comment.RepoComment) *comment.ResponderComment {
	return &comment.ResponderComment{
		RepoComment: rc,
	}
}

func TestMain(m *testing.M) {
	config.Conf = config.NewConfigTst()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()

	code := m.Run()
	os.Exit(code)
}
