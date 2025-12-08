package handlers

import (
	"chatapp/backend/dto"
	"chatapp/backend/services"
	"chatapp/backend/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IMessageHandler interface {
	CreateMessage(ctx *gin.Context, hub *ws.Hub)
	GetRoomMessages(ctx *gin.Context, hub *ws.Hub)
}

type MessageHandler struct {
	service services.IMessageService
}

func NewMessageHandler(service services.IMessageService) IMessageHandler {
	return &MessageHandler{service: service}
}

func (c *MessageHandler) CreateMessage(ctx *gin.Context, hub *ws.Hub) {
	// user, exists := ctx.Get("user")
	// if !exists {
	// ctx.AbortWithStatus(http.StatusUnauthorized)
	// return
	// }
	//
	// userId := user.(*models.User).ID

	var userId uint = 1 //仮置き

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

func (c *MessageHandler) GetRoomMessages(ctx *gin.Context, hub *ws.Hub) {

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
