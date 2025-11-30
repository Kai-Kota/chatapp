package controllers

import (
	"chatapp/backend/dto"
	"chatapp/backend/models"
	"chatapp/backend/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRoomController interface {
	CreateRoom(ctx *gin.Context)
	GetUserRooms(ctx *gin.Context)
}

type RoomController struct {
	service services.IRoomService
}

func NewRoomController(service services.IRoomService) IRoomController {
	return &RoomController{service: service}
}

func (c *RoomController) CreateRoom(ctx *gin.Context) {
	// ユーザー情報をコンテキストから取得
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// ユーザーIDを取得
	userId := user.(*models.User).ID

	// リクエストボディのバインド
	var input dto.CreateRoomInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(input)
		return
	}

	newRoom, err := c.service.Create(userId, input.Pertner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newRoom})
}

func (c *RoomController) GetUserRooms(ctx *gin.Context) {
	// ユーザー情報をコンテキストから取得
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*models.User).ID

	rooms, err := c.service.GetUserRooms(userId)
	if err != nil {
		if err.Error() == "room not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No rooms found"})
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": rooms})
}
