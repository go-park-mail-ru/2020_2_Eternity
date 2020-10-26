package comment

import (
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.Conf = config.NewConfigTst()
	config.Db = config.NewDatabase(&config.Conf.Db).Open()

	code := m.Run()
	os.Exit(code)
}

//func TestCreate(t *testing.T) {
//	c := Comment{
//		Content: "tst content",
//		PinId: 4,
//		UserId: 2,
//	}
//
//
//	assert.Nil(t, c.CreateRootComment())
//	fmt.Println(c)
//
//	assert.Nil(t, c.CreateChildComment(2))
//	fmt.Println(c)
//
//	assert.NotNil(t, c.CreateChildComment(0))
//	fmt.Println(c)
//}

//func TestGet(t *testing.T) {
//	c := Comment{}
//
//	assert.Nil(t, c.GetComment(2))
//	fmt.Println(c)
//
//	assert.NotNil(t, c.GetComment(0))
//	fmt.Println(c)
//
//}

func TestGetAll(t *testing.T) {
	c, err := GetAllComments(4)
	fmt.Println(c)
	assert.Nil(t, err)


	c, err = GetAllComments(3)
	fmt.Println(c)
	assert.NotNil(t, err)

}