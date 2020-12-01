package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/auth"
	mock_auth "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/auth/mock"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
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

func TestAuthCheck(t *testing.T) {
	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	ss, err := jwthelper.CreateJwtToken(1)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{
		Name:     config.Conf.Token.CookieName,
		Value:    ss,
		Expires:  time.Now().Add(45 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	}
	req.AddCookie(&cookie)

	c, r := gin.CreateTestContext(w)
	r.Use(AuthCheck())
	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 200)
}

func TestAuthCheckU(t *testing.T) {
	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}
	c, r := gin.CreateTestContext(w)
	r.Use(AuthCheck())
	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 401)
}

func TestAuthCheckW(t *testing.T) {

	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{
		Name:     config.Conf.Token.CookieName,
		Value:    "2134325",
		Expires:  time.Now().Add(45 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	}
	req.AddCookie(&cookie)

	c, r := gin.CreateTestContext(w)
	r.Use(AuthCheck())

	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 401)
}

func TestAuthCheckMW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	ss, err := jwthelper.CreateJwtToken(1)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{
		Name:     config.Conf.Token.CookieName,
		Value:    ss,
		Expires:  time.Now().Add(45 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	}

	req.AddCookie(&cookie)

	ac := mock_auth.NewMockAuthServiceClient(ctrl)
	mw := NewAuthMw(ac)

	c, r := gin.CreateTestContext(w)
	r.Use(mw.AuthCheck())

	ac.EXPECT().CheckCookie(gomock.Any(), &auth.Check{Cookie: ss}).Return(&auth.UserID{
		Id: 1,
	}, nil)
	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 200)
}

func TestAuthCheckMWU(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	ac := mock_auth.NewMockAuthServiceClient(ctrl)
	mw := NewAuthMw(ac)

	c, r := gin.CreateTestContext(w)

	r.Use(mw.AuthCheck())

	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 401)
}

func TestAuthCheckMwW(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	cookie := http.Cookie{
		Name:     config.Conf.Token.CookieName,
		Value:    "2134325",
		Expires:  time.Now().Add(45 * time.Minute),
		HttpOnly: true,
		Path:     "/",
	}
	req.AddCookie(&cookie)

	ac := mock_auth.NewMockAuthServiceClient(ctrl)
	mw := NewAuthMw(ac)

	c, r := gin.CreateTestContext(w)
	r.Use(mw.AuthCheck())

	ac.EXPECT().CheckCookie(gomock.Any(), &auth.Check{Cookie: "2134325"}).Return(&auth.UserID{
		Id: 1,
	}, errors.New(""))

	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 401)
}
