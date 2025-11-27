package main

import (
	"chatapp/backend/controllers"
	"chatapp/backend/infra"
	"chatapp/backend/repositories"
	"chatapp/backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	RoomRepository := repositories.NewRoomRepository(db)
	RoomService := services.NewRoomService(RoomRepository)
	RoomController := controllers.NewRoomController(RoomService)

	AuthRepository := repositories.NewAuthRepository(db)
	AuthService := services.NewAuthService(AuthRepository)
	AuthController := controllers.NewAuthController(AuthService)

	router := gin.Default()

	router.POST("/rooms", RoomController.CreateRoom)
	
	router.POST("/signup", AuthController.Signup)
	router.POST("/login", AuthController.Login)

	return router
}

func main() {
	//インフラの構築
	infra.Initialize()
	db := infra.SetupDB()
	router := SetupRouter(db)

	router.Run("localhost:8080")
}
