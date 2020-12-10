package grpc

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/auth"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type Handler struct {
	uc user.IUsecase
}

func NewHandler(uc user.IUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h Handler) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginInfo, error) {
	u, err := h.uc.GetUserByNameWithFollowers(req.Username)
	if err != nil {
		config.Lg("Auth_serv", "Login").Error(err.Error())
		return &auth.LoginInfo{
			Status: http.StatusBadRequest,
			Error:  "invalid username",
		}, errors.New("invalid username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		config.Lg("Auth_serv", "Login").Error(err.Error())
		return &auth.LoginInfo{
			Status: http.StatusBadRequest,
			Error:  "bad password",
		}, errors.New("bad password")
	}

	ss, err := jwthelper.CreateJwtToken(u.ID)
	if err != nil {
		config.Lg("Auth_serv", "Login").Error(err.Error())
		return &auth.LoginInfo{
			Status: http.StatusInternalServerError,
			Error:  "can't create token",
		}, errors.New("can't create token")
	}

	sr, err := utils.RandomUuid()
	if err != nil {
		config.Lg("Auth_serv", "Login").Error(err.Error())
		return &auth.LoginInfo{
			Status: http.StatusInternalServerError,
			Error:  "can't generate value",
		}, errors.New("can't generate value")
	}

	t, err := jwthelper.CreateCsrfToken(u.ID, sr, time.Now().Add(45*time.Minute))
	if err != nil {
		config.Lg("Auth_serv", "Login").Error(err.Error())
		return &auth.LoginInfo{
			Status: http.StatusInternalServerError,
			Error:  "can't create csrf token value",
		}, errors.New("can't create csrf token value")
	}
	return &auth.LoginInfo{
		Info: &auth.User{
			Username:    u.Username,
			Password:    u.Password,
			Avatar:      u.Avatar,
			Following:   int32(u.Following),
			Followers:   int32(u.Followers),
			Name:        u.Name,
			Description: u.Description,
			Email:       u.Email,
			Surname:     u.Surname,
		},
		Tokens: &auth.Token{
			JwtT:  ss,
			CsrfT: t,
		},
		Status: http.StatusOK,
	}, nil
}

func (h Handler) CheckCookie(ctx context.Context, check *auth.Check) (*auth.UserID, error) {
	claims := jwthelper.Claims{}
	token, err := jwthelper.ParseToken(check.Cookie, &claims)
	if err != nil {
		config.Lg("Auth_serv", "CheckCookie").Error(err.Error())
		return nil, errors.New("parse token")
	}
	if token == nil {
		return nil, errors.New("bad token")
	}
	if !token.Valid {
		return nil, errors.New("fake token")
	}
	return &auth.UserID{
		Id: int32(claims.Id),
	}, nil
}
