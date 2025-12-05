package controllers

import (
	"chatapp/backend/dto"
	"chatapp/backend/models"
	"chatapp/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IMessageController interface {
	CreateMessage(ctx *gin.Context)
	GetRoomMessages(ctx *gin.Context)
}

type MessageController struct {
	service services.IMessageService
}

func NewMessageController(service services.IMessageService) IMessageController {
	return &MessageController{service: service}
}

func (c *MessageController) CreateMessage(ctx *gin.Context) {
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

func (c *MessageController) GetRoomMessages(ctx *gin.Context) {
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
