package ws

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024

	authTime = 60 * time.Minute
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { // Note: for tests
		return true
	},
}

type Client struct {
	userId int
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
}

func NewClient(h *Hub, conn *websocket.Conn, userId int) *Client {
	return &Client{
		userId: userId,
		hub:    h,
		conn:   conn,
		send:   make(chan []byte, 256),
	}
}

func (c *Client) Register() {
	c.hub.register <- c
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		c.hub.wg.Done()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		config.Lg("ws", "readPump").Error(err.Error())
	}
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				config.Lg("ws", "readPump").Error(err.Error())
			} else {
				config.Lg("ws", "readPump").Info(err.Error())
			}
			break
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.received <- &HubMessage{
			Data:   message,
			UserId: c.userId,
		}

		config.Lg("ws", "readPump").Info("recv ", message)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	authTimer := time.NewTimer(authTime)

	defer func() {
		ticker.Stop()
		authTimer.Stop()
		c.conn.Close()
		c.hub.wg.Done()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.sendMessages(message, ok); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				config.Lg("ws", "writePump").Error("WritePing: ", err.Error())
				return
			}

		case <-authTimer.C:
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			config.Lg("ws", "writePump").Info("authTimer expired")
			return
		}
	}
}

func (c *Client) sendMessages(message []byte, ok bool) error {
	if !ok {
		c.conn.WriteMessage(websocket.CloseMessage, []byte{})
		config.Lg("ws", "sendMessages").Info("The hub closed the channel")
		return errors.New("")
	}

	if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
		config.Lg("ws", "sendMessages").Info("WriteMessage: ", err.Error())
		return err
	}

	n := len(c.send)
	for i := 0; i < n; i++ {
		if err := c.conn.WriteMessage(websocket.TextMessage, <-c.send); err != nil {
			config.Lg("ws", "sendMessages").Info("WriteMessage: ", err.Error())
			return err
		}
	}

	return nil
}
