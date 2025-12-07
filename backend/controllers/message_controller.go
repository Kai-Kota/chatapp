package controllers

import (
	"chatapp/backend/dto"
	"chatapp/backend/models"
	"chatapp/backend/services"
	"chatapp/backend/ws"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *ws.Hub
	conn *websocket.Conn
	send chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type IMessageController interface {
	ServeWs(ctx *gin.Context, hub *ws.Hub)
	CreateMessage(ctx *gin.Context, hub *ws.Hub)
	GetRoomMessages(ctx *gin.Context, hub *ws.Hub)
}

type MessageController struct {
	service services.IMessageService
}

func NewMessageController(service services.IMessageService) IMessageController {
	return &MessageController{service: service}
}

func (c *MessageController) ServeWs(ctx *gin.Context, hub *ws.Hub) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		http.Error(ctx.Writer, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
}

func (c *MessageController) CreateMessage(ctx *gin.Context, hub *ws.Hub) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*models.User).ID

	var input dto.MessageInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMessage, err := c.service.CreateMessage(userId, input.Content, input.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newMessage})
}

func (c *MessageController) GetRoomMessages(ctx *gin.Context, hub *ws.Hub) {

	var input dto.MessageGetInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messages, err := c.service.GetRoomMessages(input.RoomID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": messages})
}
