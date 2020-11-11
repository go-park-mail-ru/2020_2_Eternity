package csrf

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewTestConfig()
	return true
}()

func TestCSRFCheck(t *testing.T) {
	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	somestring := "1234142fa"

	ss, err := jwthelper.CreateCsrfToken(somestring, time.Now().Add(45*time.Minute))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-CSRF-TOKEN", ss)

	c, r := gin.CreateTestContext(w)
	r.Use(CSRFCheck())
	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 200)
}

func TestCSRFCheckW(t *testing.T) {
	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	c, r := gin.CreateTestContext(w)
	r.Use(CSRFCheck())
	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 403)
}

func TestCSRFCheckF(t *testing.T) {
	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-CSRF-TOKEN", "asdqeqwe")

	c, r := gin.CreateTestContext(w)
	r.Use(CSRFCheck())
	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 403)
}

func TestCSRFCheckE(t *testing.T) {
	w := httptest.NewRecorder()
	path := "/logout"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	somestring := "1234142fa"

	ss, err := jwthelper.CreateCsrfToken(somestring, time.Now().Add(-45*time.Minute))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-CSRF-TOKEN", ss)

	c, r := gin.CreateTestContext(w)
	r.Use(CSRFCheck())
	r.POST(path, func(c *gin.Context) {
		c.Status(200)
	})
	r.ServeHTTP(c.Writer, req)
	assert.Equal(t, c.Writer.Status(), 403)
}
