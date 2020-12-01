package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/auth"
	mock_user "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

func TestHandler_Login(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_user.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	testUser := &auth.LoginReq{Username: "21savage",
		Password: "12345678",
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), 7)
	if err != nil {
		log.Fatal(err)
	}

	respUser := &domain.User{
		Username: "21savage",
		Password: string(hash),
	}

	uc.EXPECT().GetUserByNameWithFollowers(gomock.Any()).Return(respUser, nil)

	info, err := handler.Login(context.Background(), &auth.LoginReq{Username: "21savage",
		Password: "12345678"},
	)
	assert.NoError(t, err)

	assert.Equal(t, int32(200), info.Status)
}

func TestHandler_LoginF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_user.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	respUser := &domain.User{
		Username: "21savage",
		Password: "12123241",
	}

	uc.EXPECT().GetUserByNameWithFollowers(gomock.Any()).Return(respUser, nil)

	info, err := handler.Login(context.Background(), &auth.LoginReq{Username: "21savage",
		Password: "12345678"},
	)
	assert.Error(t, err)

	assert.Equal(t, int32(400), info.Status)
}

func TestHandler_CheckCookie(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_user.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	ss, err := jwthelper.CreateJwtToken(1)
	if err != nil {
		log.Fatal(err)
	}
	check := &auth.Check{
		Cookie: ss,
	}

	id, err := handler.CheckCookie(context.Background(), check)
	assert.NoError(t, err)
	assert.Equal(t, 1, int(id.Id))
}

func TestHandler_CheckCookieF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock_user.NewMockIUsecase(ctrl)

	handler := NewHandler(uc)

	check := &auth.Check{
		Cookie: "12345",
	}

	_, err := handler.CheckCookie(context.Background(), check)
	assert.Error(t, err)
}
