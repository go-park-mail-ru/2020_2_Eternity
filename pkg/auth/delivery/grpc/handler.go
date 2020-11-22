package grpc

import (
	"context"
	"errors"
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
		return &auth.LoginInfo{
			Status: http.StatusBadRequest,
			Error:  "invalid username",
		}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return &auth.LoginInfo{
			Status: http.StatusBadRequest,
			Error:  "bad password",
		}, err
	}

	ss, err := jwthelper.CreateJwtToken(u.ID)
	if err != nil {
		return &auth.LoginInfo{
			Status: http.StatusInternalServerError,
			Error:  "can't create token",
		}, err
	}

	sr, err := utils.RandomUuid()
	if err != nil {
		return &auth.LoginInfo{
			Status: http.StatusInternalServerError,
			Error:  "can't generate value",
		}, err
	}

	t, err := jwthelper.CreateCsrfToken(u.ID, sr, time.Now().Add(45*time.Minute))
	if err != nil {
		return &auth.LoginInfo{
			Status: http.StatusInternalServerError,
			Error:  "can't create csrf token value",
		}, err
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
		return nil, err
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
