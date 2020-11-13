package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
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
