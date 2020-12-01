package grpc

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

func TestHandler_Login(t *testing.T) {
	/*t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_user.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	srv := mock_auth.NewMockAuthServiceServer(ctrl)
	*/
}
