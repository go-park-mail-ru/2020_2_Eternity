package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	mock_database "github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database/mock"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	mock_user "github.com/go-park-mail-ru/2020_2_Eternity/pkg/user/mock"
	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

var p = bluemonday.UGCPolicy()

func TestDelivery_SignUpSuccess(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/signup"

	testValidUser := api.SignUp{
		Username: "21savage",
		Email:    "kae@email.com",
		Password: "12345678",
	}

	respUser := &domain.User{}

	body, err := json.Marshal(testValidUser)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	userMockUsecase.EXPECT().CreateUser(gomock.Any()).Return(respUser, nil)

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)

	r.POST(path, userHandler.SignUp)
	r.ServeHTTP(c.Writer, req)

	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_SignUpValidP(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/signup"

	testUser := api.SignUp{
		Username: "21savage",
		Email:    "kaeemail.com",
		Password: "1234578",
	}
	body, err := json.Marshal(testUser)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	c, r := gin.CreateTestContext(w)
	r.POST(path, userHandler.SignUp)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_LoginF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/login"

	testUser := api.Login{
		Username: "21savage",
		Password: "1234578",
	}
	body, err := json.Marshal(testUser)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}
	respUser := &domain.User{}
	userMockUsecase.EXPECT().GetUserByNameWithFollowers(gomock.Any()).Return(respUser, nil)

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	_, r := gin.CreateTestContext(w)
	r.POST(path, userHandler.Login)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestDelivery_Login(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/login"

	testUser := api.Login{
		Username: "21savage",
		Password: "1234578",
	}
	body, err := json.Marshal(testUser)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(testUser.Password), config.Conf.Token.Value)
	if err != nil {
		log.Fatal(err)
	}

	respUser := &domain.User{
		Username: "21savage",
		Password: string(hash),
	}
	userMockUsecase.EXPECT().GetUserByNameWithFollowers(gomock.Any()).Return(respUser, nil)

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	_, r := gin.CreateTestContext(w)
	r.POST(path, userHandler.Login)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDelivery_LogoutF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}
	_, r := gin.CreateTestContext(w)
	r.POST(path, userHandler.Logout)
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestDelivery_LogoutS(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{
		Name:     config.Conf.Token.CookieName,
		Value:    "12345",
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	req.AddCookie(&cookie)

	c, r := gin.CreateTestContext(w)

	r.POST(path, userHandler.Logout)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func mid() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("info", 1)
		c.Next()
	}
}

func TestDelivery_UpdateS(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile"

	testUser := api.UpdateUser{
		Username: "21savage",
	}
	body, err := json.Marshal(testUser)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	respUser := &domain.User{}
	userMockUsecase.EXPECT().UpdateUser(1, gomock.Any()).Return(respUser, nil)

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)

	r.Use(mid())
	r.PUT(path, userHandler.UpdateUser)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_UpdateC(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile"

	testUser := api.UpdateUser{
		Username: "21savage",
	}
	body, err := json.Marshal(testUser)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	userMockUsecase.EXPECT().UpdateUser(1, gomock.Any()).Return(nil, errors.New(""))

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)

	r.Use(mid())
	r.PUT(path, userHandler.UpdateUser)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 422, c.Writer.Status())
}

func TestDelivery_UpdateUserUnAuth(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile"
	req, _ := http.NewRequest("PUT", path, nil)
	c, r := gin.CreateTestContext(w)
	r.PUT(path, userHandler.UpdateUser)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 401, c.Writer.Status())
}

func TestDelivery_UpdateUserFail(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile"

	testUser := api.UpdateUser{
		Username: "21savage4^&",
	}
	body, err := json.Marshal(testUser)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)

	r.Use(mid())
	r.PUT(path, userHandler.UpdateUser)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_UpdatePasswordF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile"

	testPswd := api.UpdatePassword{
		OldPassword: "21savage4^&",
		NewPassword: "1231",
	}
	body, err := json.Marshal(testPswd)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.PUT(path, userHandler.UpdatePassword)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_UpdatePassword(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile/password"

	testPswd := api.UpdatePassword{
		OldPassword: "12345678",
		NewPassword: "123145678",
	}
	body, err := json.Marshal(testPswd)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(testPswd.OldPassword), config.Conf.Token.Value)
	if err != nil {
		log.Fatal(err)
	}
	u := domain.User{
		Password: string(hash),
	}

	userMockUsecase.EXPECT().GetUser(gomock.Any()).Return(&u, nil)
	userMockUsecase.EXPECT().UpdatePassword(gomock.Any(), gomock.Any()).Return(nil)

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.PUT(path, userHandler.UpdatePassword)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_UpdatePasswordU(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile/password"

	testPswd := api.UpdatePassword{
		OldPassword: "12345678",
		NewPassword: "123145678",
	}
	body, err := json.Marshal(testPswd)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(testPswd.OldPassword), config.Conf.Token.Value)
	if err != nil {
		log.Fatal(err)
	}
	u := domain.User{
		Password: string(hash),
	}

	userMockUsecase.EXPECT().GetUser(gomock.Any()).Return(&u, nil)
	userMockUsecase.EXPECT().UpdatePassword(gomock.Any(), gomock.Any()).Return(errors.New(""))

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.PUT(path, userHandler.UpdatePassword)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 500, c.Writer.Status())
}

func TestDelivery_UpdatePasswordG(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile/password"

	testPswd := api.UpdatePassword{
		OldPassword: "12345678",
		NewPassword: "123145678",
	}
	body, err := json.Marshal(testPswd)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	u := domain.User{}

	userMockUsecase.EXPECT().GetUser(gomock.Any()).Return(&u, errors.New(""))

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.PUT(path, userHandler.UpdatePassword)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_UpdatePasswordAuth(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile/password"

	testPswd := api.UpdatePassword{
		OldPassword: "12345678",
		NewPassword: "123145678",
	}
	body, err := json.Marshal(testPswd)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}
	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))
	c, r := gin.CreateTestContext(w)
	r.PUT(path, userHandler.UpdatePassword)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 401, c.Writer.Status())
}

func TestDelivery_UpdatePasswordW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()

	path := "/user/profile/password"

	testPswd := api.UpdatePassword{
		OldPassword: "12345678",
		NewPassword: "123145678",
	}
	body, err := json.Marshal(testPswd)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	u := domain.User{}

	userMockUsecase.EXPECT().GetUser(gomock.Any()).Return(&u, nil)

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.PUT(path, userHandler.UpdatePassword)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_GetProfileUnAuth(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()
	path := "/user/profile"
	req, _ := http.NewRequest("GET", path, nil)

	c, r := gin.CreateTestContext(w)
	r.GET(path, userHandler.GetProfile)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 401, c.Writer.Status())
}

func TestDelivery_GetProfile(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()
	path := "/user/profile"
	req, _ := http.NewRequest("GET", path, nil)

	u := &domain.User{}
	userMockUsecase.EXPECT().GetUser(1).Return(u, nil)

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.GET(path, userHandler.GetProfile)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_GetProfileFail(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()
	path := "/user/profile"
	req, _ := http.NewRequest("GET", path, nil)

	u := &domain.User{}
	userMockUsecase.EXPECT().GetUser(gomock.Any()).Return(u, errors.New("fake"))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.GET(path, userHandler.GetProfile)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_GetAvatarFail(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	w := httptest.NewRecorder()
	path := "/image/avatar/fff"
	req, _ := http.NewRequest("GET", path, nil)

	c, r := gin.CreateTestContext(w)
	r.GET("/image/avatar/:file", userHandler.GetAvatar)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_Follow(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21savage",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	u := &domain.User{Username: "21savage"}

	userMockUsecase.EXPECT().GetUserByName(u.Username).Return(u, nil)
	userMockUsecase.EXPECT().Follow(1, gomock.Any()).Return(nil)
	w := httptest.NewRecorder()
	path := "/follow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, userHandler.Follow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_FollowE(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21savage",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	u := &domain.User{Username: "21savage"}

	userMockUsecase.EXPECT().GetUserByName(u.Username).Return(u, nil)
	userMockUsecase.EXPECT().Follow(1, gomock.Any()).Return(errors.New(""))
	w := httptest.NewRecorder()
	path := "/follow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, userHandler.Follow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_FollowV(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	w := httptest.NewRecorder()
	path := "/follow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, userHandler.Follow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_FollowF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21savage",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	w := httptest.NewRecorder()
	path := "/follow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.POST(path, userHandler.Follow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 401, c.Writer.Status())
}

func TestDelivery_UnFollow(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21savage",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	u := &domain.User{Username: "21savage"}

	userMockUsecase.EXPECT().GetUserByName(u.Username).Return(u, nil)
	userMockUsecase.EXPECT().UnFollow(1, gomock.Any()).Return(nil)
	w := httptest.NewRecorder()
	path := "/unfollow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, userHandler.Unfollow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_UnFollowE(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21savage",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	u := &domain.User{Username: "21savage"}

	userMockUsecase.EXPECT().GetUserByName(u.Username).Return(u, errors.New(""))
	w := httptest.NewRecorder()
	path := "/unfollow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, userHandler.Unfollow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_UnFollowW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21savage",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	u := &domain.User{Username: "21savage"}

	userMockUsecase.EXPECT().GetUserByName(u.Username).Return(u, nil)
	userMockUsecase.EXPECT().UnFollow(1, gomock.Any()).Return(errors.New(""))
	w := httptest.NewRecorder()
	path := "/unfollow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.Use(mid())
	r.POST(path, userHandler.Unfollow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_UnFollowF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	f := api.UserAct{
		Username: "21savage",
	}

	body, err := json.Marshal(f)
	if err != nil {
		log.Fatal("cant marshal")
		return
	}

	w := httptest.NewRecorder()
	path := "/unfollow"
	req, _ := http.NewRequest("POST", path, bytes.NewReader(body))

	c, r := gin.CreateTestContext(w)
	r.POST(path, userHandler.Unfollow)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 401, c.Writer.Status())
}

func TestDelivery_UserPageS(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	u := &domain.User{Username: "21savage"}

	userMockUsecase.EXPECT().GetUserByNameWithFollowers(u.Username).Return(u, nil)

	w := httptest.NewRecorder()
	path := "/userpage/21savage"
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.GET("/userpage/:username", userHandler.GetUserPage)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_UserPageF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	u := &domain.User{Username: "21"}

	userMockUsecase.EXPECT().GetUserByNameWithFollowers(u.Username).Return(u, errors.New(""))

	w := httptest.NewRecorder()
	path := "/userpage/21"
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.GET("/userpage/:username", userHandler.GetUserPage)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_GetFollowersS(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	var users []domain.User

	userMockUsecase.EXPECT().GetFollowers("21savage").Return(users, nil)

	w := httptest.NewRecorder()
	path := "/followers/21savage"
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.GET("/followers/:username", userHandler.GetFollowers)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_GetFollowersF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	var users []domain.User

	userMockUsecase.EXPECT().GetFollowers("21").Return(users, errors.New(""))

	w := httptest.NewRecorder()
	path := "/followers/21"
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.GET("/followers/:username", userHandler.GetFollowers)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestDelivery_GetFollowingS(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	var users []domain.User

	userMockUsecase.EXPECT().GetFollowing("21savage").Return(users, nil)

	w := httptest.NewRecorder()
	path := "/following/21savage"
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.GET("/following/:username", userHandler.GetFollowing)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestDelivery_GetFollowingF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)

	var users []domain.User

	userMockUsecase.EXPECT().GetFollowing("21").Return(users, errors.New(""))

	w := httptest.NewRecorder()
	path := "/following/21"
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.GET("/following/:username", userHandler.GetFollowing)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 400, c.Writer.Status())
}

func TestHandler_IsFollowing(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)
	w := httptest.NewRecorder()

	path := "/isfollowing/21savage"

	userMockUsecase.EXPECT().IsFollowing(1, "21savage").Return(nil)
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.Use(mid())

	r.GET("/isfollowing/:username", userHandler.IsFollowing)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_IsFollowingF(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)
	w := httptest.NewRecorder()

	path := "/isfollowing/21"

	userMockUsecase.EXPECT().IsFollowing(1, "21").Return(errors.New(""))
	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)
	r.Use(mid())

	r.GET("/isfollowing/:username", userHandler.IsFollowing)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestHandler_IsFollowingU(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userMockUsecase := mock_user.NewMockIUsecase(ctrl)
	userHandler := NewHandler(userMockUsecase, p)
	w := httptest.NewRecorder()

	path := "/isfollowing/21savage"

	req, _ := http.NewRequest("GET", path, nil)
	c, r := gin.CreateTestContext(w)

	r.GET("/isfollowing/:username", userHandler.IsFollowing)
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, 401, c.Writer.Status())
}


func TestCreateRoutes(t *testing.T) {
	mockCtr := gomock.NewController(t)
	defer mockCtr.Finish()

	writerResp := httptest.NewRecorder()
	_, r := gin.CreateTestContext(writerResp)

	mockDatabase := mock_database.NewMockIDbConn(mockCtr)
	AddUserRoutes(r, mockDatabase, bluemonday.NewPolicy())
}