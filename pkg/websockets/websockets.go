package websockets

import (
	"container/list"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type WebSocketPool struct {
	upgrader websocket.Upgrader
	clients  map[int]*list.List //*websocket.Conn
	mutex    *sync.Mutex
}

func NewPool() *WebSocketPool {
	return &WebSocketPool{
		clients: make(map[int]*list.List, 0),
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
	id, ok := c.Get("id") // надо выставлять после проверки аутентификации, т.е использовать с auth mw
	if !ok {
		log.Println("ID doesnt set")
		//return
	}

	wsPool.mutex.Lock()
	conn, err := wsPool.upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
	if err != nil {
		log.Println(err)
		return
	}
	id = 1
	if _, ok := wsPool.clients[id.(int)]; !ok {
		wsPool.clients[id.(int)] = list.New()
	}

	wsPool.clients[id.(int)].PushFront(conn)
	//wsPool.clients[id.(int)].PushFront() = //append(wsPool.clients[id.(int)], conn) // сюда не 1, а id из get
	wsPool.mutex.Unlock()
}

func (wsPool *WebSocketPool) Send(uid int, message []byte) {
	l, ok := wsPool.clients[uid]
	if !ok {
		log.Println("Empty")
		return
	}
	for e := l.Front(); e != nil; e = e.Next() {
		if err := e.Value.(*websocket.Conn).WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			log.Println(err)
			err = e.Value.(*websocket.Conn).Close()
			if err != nil {
				log.Println(err)
			}
			wsPool.mutex.Lock()
			l.Remove(e)
			wsPool.mutex.Unlock()
			continue
		}

		if err := e.Value.(*websocket.Conn).WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println(err)
			err = e.Value.(*websocket.Conn).Close()
			if err != nil {
				log.Println(err)
			}
			wsPool.mutex.Lock()
			l.Remove(e)
			wsPool.mutex.Unlock()
		}
	}

	if _, ok := wsPool.clients[uid]; ok && wsPool.clients[uid].Len() == 0 {
		wsPool.mutex.Lock()
		delete(wsPool.clients, uid)
		// удалить id клиента из мапы под мьютексом
		log.Println("DELETE NEED")
		wsPool.mutex.Unlock()
	}
}
