package pintests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"math"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var (
	ts = httptest.NewServer(setupServer())

	u = user.User{
		ID:        math.MaxInt32 - 10,
		Username:  "username23456789876543234567",
		Email:     "email235462345643u526453446346253",
		Password:  "123",
		BirthDate: time.Now(),
	}
)

func setupServer() *gin.Engine {
	r := gin.Default()
	r.POST("/user/pin", createClaims, pin.CreatePin)
	r.GET("/user/pin", createClaims, pin.GetPin)

	return r
}

func createClaims(c *gin.Context) {
	claims := jwthelper.Claims{
		Id: u.ID,
	}
	c.Set("info", claims.Id)
	c.Next()
}

func TestMain(m *testing.M) {
	code := m.Run()
	ts.Close()
	os.Exit(code)
}
