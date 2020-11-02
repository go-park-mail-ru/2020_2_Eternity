package websockets

import (
	"github.com/gin-gonic/gin"
	"log"
)

func TestMwWs(ws *WebSocketPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status, ok := c.Get("status")
		if !ok {
			log.Println("STATUS")
			return
		}
		uid, ok := c.Get("follow_id")
		if !ok {
			log.Println("ID IS NOT SET")
			return
		}
		ws.Send(uid.(int), []byte("kto-to podpisalsya grats!"))
		log.Println(status, uid)
	}
}
