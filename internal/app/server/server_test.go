package server

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	mock_database "github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()


func TestNew(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	mockDatabase := mock_database.NewMockIDbConn(mockCtr)

	New(config.Conf, mockDatabase)
}

