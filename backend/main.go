package main

import (
	"chatapp/backend/controllers"
	"chatapp/backend/infra"
	middleware "chatapp/backend/middlewares"
	"chatapp/backend/repositories"
	"chatapp/backend/services"
	ws "chatapp/backend/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	RoomRepository := repositories.NewRoomRepository(db)
	RoomService := services.NewRoomService(RoomRepository)
	RoomController := controllers.NewRoomController(RoomService)

	MessageRepository := repositories.NewMessageRepository(db)
	MessageService := services.NewMessageService(MessageRepository)
	MessageHandler := ws.NewMessageHandler(MessageService)

	AuthRepository := repositories.NewAuthRepository(db)
	AuthService := services.NewAuthService(AuthRepository)
	AuthController := controllers.NewAuthController(AuthService)

	hub := ws.NewHub()
	go hub.Run()

	router := gin.Default()

	//CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Upgrade", "Connection"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,	
		MaxAge:           12 * 60 * 60,
	}))
	routerWithAuth := router.Group("/user", middleware.AuthMiddleware(AuthService))
	// WebSocket ハンドラ登録
	wsrouter := router.Group("/ws")
	wsrouter.GET("/messages", func(ctx *gin.Context) {
		MessageHandler.ServeWs(hub, ctx)
	})

	router.POST("/signup", AuthController.Signup)
	router.POST("/login", AuthController.Login)

	routerWithAuth.POST("/rooms", RoomController.CreateRoom)
	routerWithAuth.GET("/rooms", RoomController.GetUserRooms)

	return router
}

func main() {
	//インフラの構築
	infra.Initialize()
	db := infra.SetupDB()
	router := SetupRouter(db)

	router.Run("localhost:8080")
}
