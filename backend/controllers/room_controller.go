package controllers

import (
	"chatapp/backend/dto"
	"chatapp/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRoomController interface {
	CreateRoom(ctx *gin.Context)
}

type RoomController struct {
	service services.IRoomService
}

func NewRoomController(service services.IRoomService) IRoomController {
	return &RoomController{service: service}
}

func (c *RoomController) CreateRoom(ctx *gin.Context) {
	var input dto.CreateRoomInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRoom, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newRoom})
}
