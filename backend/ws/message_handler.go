package ws

import (
	"chatapp/backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IMessageHandler interface {
	serveWs(c *MessageClient)
}

type MessageHandler struct {
	service services.IMessageService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewMessageHandler(service services.IMessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func (c *MessageHandler) ServeWs(hub *Hub, ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	client := newMessageClient(hub, conn)
	client.hub.register <- client

	go client.CreateMessage(c.service)
	go client.GetRoomMessages(c.service)
}
