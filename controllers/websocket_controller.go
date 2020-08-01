package controllers

import (
	"fmt"
	"sync/atomic"
	"simpleapi/models"
	"simpleapi/services"

    "github.com/kataras/iris/v12/websocket"
)

type WebsocketController struct {
	*websocket.NSConn `stateless:"true"`
	Services services.MessageService
	Namespace string
}

var visits uint64

func increment() uint64 {
	return atomic.AddUint64(&visits, 1)
}

func decrement() uint64 {
	return atomic.AddUint64(&visits, ^uint64(0))
}

func (c *WebsocketController) OnNamespaceDisconnect(msg websocket.Message) error {
	// visits--
	newCount := decrement()
	// This will call the "OnVisit" event on all clients, except the current one,
	// (it can't because it's left but for any case use this type of design)
	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "OnVisit",
		Body:      []byte(fmt.Sprintf("%d", newCount)),
	})

	return nil
}

func (c *WebsocketController) OnNamespaceConnected(msg websocket.Message) error {
	// visits++
	newCount := increment()

	// This will call the "OnVisit" event on all clients, including the current one,
	// with the 'newCount' variable.
	c.Conn.Server().Broadcast(nil, websocket.Message{
		Namespace: msg.Namespace,
		Event:     "OnVisit",
		Body:      []byte(fmt.Sprintf("%d", newCount)),
	})

	return nil
}

func (c *WebsocketController) OnChat(msg websocket.Message) error {
	ctx := websocket.GetContext(c.Conn)

	var dataModels models.Message
	dataModels.Message = string(msg.Body)
	_, err := c.Services.CreateMessage(dataModels, ctx)
	if err != nil {
		return err
	}

	c.Conn.Server().Broadcast(c, msg)

	return nil
}