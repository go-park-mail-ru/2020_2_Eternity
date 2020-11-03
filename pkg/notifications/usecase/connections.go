package usecase

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"sync"
	"time"
)

const (
	connExpireTime = 5
)

type Connect struct {
	cn  *websocket.Conn
	expire time.Time
}

type Connections struct {
	clients  map[int][]Connect
	mux sync.Mutex
}

func NewConnections() *Connections {
	return &Connections{
		clients:  make(map[int][]Connect),
		mux: sync.Mutex{},
	}
}

func (c *Connections) Add(conn *websocket.Conn, userId int) {
	c.mux.Lock()
	c.clients[userId] = append(c.clients[userId], Connect{
		cn: conn,
		expire: time.Now().Add(connExpireTime * time.Minute),
	})
	c.mux.Unlock()
}

func (c *Connections) SendData(userId int, data []byte) error {
	c.mux.Lock()
	connections, ok := c.clients[userId]
	c.mux.Unlock()

	if !ok {
		return errors.New("Connection not exists")
	}

	newConnections := []Connect{}
	for _, conn := range connections {
		if err := sendOne(&conn, data); err != nil {
			conn.cn.Close()
		} else {
			newConnections = append(newConnections, conn)
		}
	}

	if len(newConnections) == 0 {
		c.mux.Lock()
		delete(c.clients, userId)
		c.mux.Unlock()
	} else {
		c.mux.Lock()
		c.clients[userId] = newConnections
		c.mux.Unlock()
	}

	return nil
}

func sendOne(conn *Connect, data []byte) error {
	if conn.expire.Sub(time.Now()) < 0 {
		return errors.New("Connect expired")
	}

	if err := conn.cn.WriteMessage(websocket.TextMessage, data); err != nil {
		return err
	}

	return nil
}
