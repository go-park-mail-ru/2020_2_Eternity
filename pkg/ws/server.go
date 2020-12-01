package ws

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	domainWs "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/ws"
	"net/http"
)


type IServer interface{
	SetHandler(msgType string, handlers ...func(c *Context))
	Run()
	Stop()
	SendMessage(msg *domainWs.MessageResp)
	CloseConnection(userId int)
	RegisterClient(w http.ResponseWriter, req *http.Request, userId int) error
}


type Context struct {
	Req  domainWs.MessageReq
	Resp []domainWs.MessageResp
	vars map[string]interface{}
	handlers []func(c *Context)
	currentHandler int
}

func NewContext(req domainWs.MessageReq, handlers []func(c *Context)) *Context {
	return &Context{
		Req: req,
		handlers: handlers,
		currentHandler: 0,
	}
}


func (c *Context) Next() {
	if c.currentHandler >= len(c.handlers) {
		config.Lg("ws", "Next").
			Errorf("currentHandler index = %d out of range = %d", c.currentHandler, len(c.handlers))
		return
	}

	curr := c.currentHandler
	c.currentHandler += 1

	c.handlers[curr](c)
}

func (c *Context) Set(varName string, variable interface{}) {
	if c.vars == nil {
		c.vars = make(map[string]interface{})
	}

	c.vars[varName] = variable
}

func (c *Context) Get(varName string) (interface{}, error) {
	if c.vars == nil {
		return nil, errors.New("Map not created")
	}

	value, ok := c.vars[varName]
	if !ok {
		return nil, errors.New("Value not exists")
	}

	return value, nil
}


func (c *Context) AbortWithStatus(msgType string, userId, status int) {
	c.Resp = append(c.Resp, domainWs.MessageResp{
		UserId: userId,
		Type: msgType,
		Status: status,
	})
}

func (c *Context) AddResponse(msg interface{}, msgType string, userId, status int) {
	data, err := json.Marshal(msg)
	if err != nil {
		config.Lg("ws_server", "AddResponse").
			Error("Marshal: ", err.Error())

		c.AbortWithStatus(msgType, userId, status)
		return
	}

	c.Resp = append(c.Resp, domainWs.MessageResp{
		UserId: userId,
		Type: msgType,
		Status: status,
		Data: data,
	})
}




// Router
type Router struct {
	routes map[string][]func(c *Context)
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string][]func(c *Context)),
	}
}

func (r *Router) SetHandler(msgType string, handlers ...func(c *Context))  {
	for _, h := range handlers {
		r.routes[msgType] = append(r.routes[msgType], h)
	}
}



// Server
type Server struct {
	hub      IHub
	*Router
	received chan *HubMessage
	toSend   chan *HubMessage
}

func NewServer() *Server {
	h := NewHub()
	return &Server{
		hub: h,
		Router: NewRouter(),
		received: h.GetRecvChanel(),
		toSend: h.GetSendChanel(),
	}
}

func (s *Server) Run() {
	s.hub.Run()
	go s.handleMessages()
}

func (s *Server) Stop() {
	s.hub.Stop()
}

func (s *Server) SendMessage(msg *domainWs.MessageResp) {
	defer func() {
		if recover() != nil {
			config.Lg("ws_server", "SendMessage").
				Error("Recover ...")
		}
	}()

	data, err := json.Marshal(msg)
	if err != nil {
		config.Lg("ws_server", "SendMessage").
			Error("Marshal: ", err.Error())
		return
	}

	s.toSend <- &HubMessage{
		UserId: msg.UserId,
		Data: data,
	}
}

func (s *Server) CloseConnection(userId int) {
	s.hub.CloseConnection(userId)
}

func (s *Server) RegisterClient(w http.ResponseWriter, req *http.Request, userId int) error {
	conn, err := Upgrader.Upgrade(w, req, nil)
	if err != nil {
		config.Lg("ws", "RegisterClient").Error(err.Error())
		return err
	}

	s.hub.RegisterClient(conn, userId)

	return nil
}


func (s *Server) handleMessages() {
	for m := range s.received {

		//For tests
		//s.toSend<- &HubMessage{2, m.Data}
		//continue
		//For tests


		msg := domainWs.MessageReq{}

		if err := json.Unmarshal(m.Data, &msg); err != nil {
			config.Lg("ws", "handleMessages").
				Error("Unmarshal: ", err.Error())
			continue
		}
		msg.UserId = m.UserId

		//c := Context{Req: msg}

		handlers, ok := s.routes[msg.Type]
		if !ok {
			config.Lg("ws", "handleMessages").
				Errorf("Handler for type '%s' not exists", msg.Type)
			continue
		}

		//handler(&c)

		c := NewContext(msg, handlers)
		c.Next()

		for _, resp := range c.Resp {
			s.SendMessage(&resp)
		}
	}
}


