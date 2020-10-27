package websockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WebSocketPool struct {
	upgrader websocket.Upgrader
	clients  map[int][]*websocket.Conn
	mutex    *sync.Mutex
}

func NewPool() *WebSocketPool {
	return &WebSocketPool{
		clients: make(map[int][]*websocket.Conn, 0),
		upgrader: websocket.Upgrader{
			WriteBufferSize: 1024,
			ReadBufferSize:  1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		mutex: &sync.Mutex{},
	}
}

func (wsPool *WebSocketPool) Add(c *gin.Context) {
	_, ok := c.Get("id") // надо выставлять после проверки аутентификации, т.е использовать с auth mw
	if !ok {
		log.Println("SAMO SOBOY BEZ MW ZHE")
	}

	wsPool.mutex.Lock()
	conn, err := wsPool.upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		log.Println(err)
		return
	}
	wsPool.clients[1] = append(wsPool.clients[1], conn) // сюда не 1, а id из get
	wsPool.mutex.Unlock()
}

func (wsPool *WebSocketPool) Send(uid int, message []byte) {
	for _, client := range wsPool.clients[uid] {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
		}
		log.Println(err)
	}
	log.Println(" len is ", len(wsPool.clients))
}
