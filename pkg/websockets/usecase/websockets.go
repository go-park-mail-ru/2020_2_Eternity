package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

const (
	WorkersAmount  = 2
	queueSize = 1000
)



type WebSocketPool struct {
	Upgrader websocket.Upgrader
	connects  *Connections
	noteQueue chan domain.Notification
	cancel context.CancelFunc
	wg sync.WaitGroup
}

func NewPool() *WebSocketPool {
	return &WebSocketPool{
		Upgrader: websocket.Upgrader{
			WriteBufferSize: 1024,
			ReadBufferSize:  1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		connects:  NewConnections(),
		noteQueue: make(chan domain.Notification, queueSize),
		wg: sync.WaitGroup{},
	}
}

func (wsPool *WebSocketPool) AddConnection(conn *websocket.Conn, userId int) {
	wsPool.connects.Add(conn, userId)
}

func (wsPool *WebSocketPool) AddNotification(note *domain.Notification) {
	wsPool.noteQueue <- *note
}


//func (wsPool *WebSocketPool) Add(userId int) {
//
//	wsPool.mutex.Lock()
//	conn, err := wsPool.upgrader.Upgrade(c.Writer, c.Request, c.Writer.Header())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	wsPool.clients[userId] =  conn // сюда не 1, а id из get
//	wsPool.mutex.Unlock()
//}

//func (wsPool *WebSocketPool) Send(uid int, message []byte) {
//	for _, client := range wsPool.clients[uid] {
//		err := client.WriteMessage(websocket.TextMessage, message)
//		if err != nil {
//			log.Println(err)
//		}
//		log.Println(err)
//	}
//	log.Println(" len is ", len(wsPool.clients))
//}

func (wsPool *WebSocketPool) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	wsPool.cancel = cancel

	for i := 0; i < WorkersAmount; i++ {
		go func(i int) {
			wsPool.wg.Add(1)

			var note domain.Notification
			select {
			case <- ctx.Done():
				config.Lg("websockets", fmt.Sprintf("Worker#%d", i)).Info("Goroutine exit")
				wsPool.wg.Done()
				return
			case note = <- wsPool.noteQueue:
				wsPool.notify(&note)
			}
		}(i)
	}
}

func (wsPool *WebSocketPool) notify(note *domain.Notification) {
	data, err := json.Marshal(note)
	if err != nil {
		config.Lg("websockets", "notify").Error(err.Error())
		return
	}

	if err := wsPool.connects.SendData(note.ToUserId, data); err != nil {
		config.Lg("websockets", "notify").Error(err.Error())
		return
	}
}

func (wsPool *WebSocketPool) Stop() {
	// TODO (Pavel S) close connections
	if wsPool.cancel == nil {
		config.Lg("websockets", "Stop").Error("Nothing to exit")
		return
	}

	config.Lg("websockets", "Stop").Info("Terminating workers...")
	wsPool.cancel()
	wsPool.wg.Wait()
}