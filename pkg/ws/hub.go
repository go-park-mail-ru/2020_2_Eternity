package ws

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/gorilla/websocket"
	"sync"
)

type IHub interface {
	Run()
	Stop()
	GetRecvChanel() chan *HubMessage
	GetSendChanel() chan *HubMessage
	CloseConnection(userId int)
	RegisterClient(conn *websocket.Conn, userId int)
}

type HubMessage struct {
	UserId int
	Data   []byte
}

type Hub struct {
	clients map[int][]*Client
	mux     sync.Mutex
	wg      sync.WaitGroup

	toSend     chan *HubMessage
	received   chan *HubMessage
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int][]*Client),
		mux:     sync.Mutex{},
		wg:      sync.WaitGroup{},

		toSend:     make(chan *HubMessage, 256),
		received:   make(chan *HubMessage, 256),
		register:   make(chan *Client, 256),
		unregister: make(chan *Client, 256),
	}
}


func (h *Hub) GetRecvChanel() chan *HubMessage {
	return h.received
}

func (h *Hub) GetSendChanel() chan *HubMessage {
	return h.toSend
}

func (h *Hub) RegisterClient(conn *websocket.Conn, userId int) {
	client := NewClient(h, conn, userId)
	client.Register()
}

func (h *Hub) registerClient() {
	for c := range h.register {

		h.mux.Lock()
		h.clients[c.userId] = append(h.clients[c.userId], c)
		h.mux.Unlock()

		go c.readPump()
		go c.writePump()

		h.wg.Add(2)
	}
}

func (h *Hub) unregisterClient() {
	for c := range h.unregister {
		safeClose(c.send)

		h.mux.Lock()

		clients, ok := h.clients[c.userId]

		config.Lg("ws", "unregisterClient").Debug("before ", h.clients)

		if ok {
			for i, client := range clients {
				if client == c {
					clients[i] = clients[len(clients)-1]
					clients = clients[:len(clients)-1]

					if len(clients) == 0 {
						delete(h.clients, c.userId)
					} else {
						h.clients[c.userId] = clients
					}

					break
				}
			}
		}

		config.Lg("ws", "unregisterClient").Debug("after ", h.clients)

		h.mux.Unlock()
	}
}

func (h *Hub) sendMsgWorker() {
	for msg := range h.toSend {
		h.mux.Lock()

		clients, ok := h.clients[msg.UserId]
		if ok {
			for _, c := range clients {
				c.send <- msg.Data
			}

			config.Lg("ws", "sendMsgWorker").Debug("send ", msg)
		} else {
			config.Lg("ws", "sendMsgWorker").Debug("No clients found")
		}

		h.mux.Unlock()
	}
}

//func (h *Hub) receiveMsgWorker() {
//	for msg := range h.received {
//		fmt.Println(msg)
//
//		config.Lg("ws", "receiveMsgWorker").Debug(msg)
//
//		// msgs := grpc.SomeFunc(msg)
//		// send responses back
//
//		h.toSend <- &HubMessage{2, msg.Data}
//	}
//}

func (h *Hub) Run() {
	go h.registerClient()
	go h.unregisterClient()
	go h.sendMsgWorker()
	//go h.receiveMsgWorker()
}

func (h *Hub) Stop() {
	h.mux.Lock()

	for _, con := range h.clients {
		for _, c := range con {
			close(c.send)
		}
	}

	h.mux.Unlock()

	h.wg.Wait()

	close(h.register)
	close(h.unregister)
	close(h.toSend)
	close(h.received)
}

func (h *Hub) CloseConnection(userId int) {
	h.mux.Lock()
	clients, ok := h.clients[userId]
	h.mux.Unlock()

	if ok {
		for _, c := range clients {
			close(c.send)
		}
	}
}


func safeClose(ch chan []byte) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	close(ch)
	return true
}


