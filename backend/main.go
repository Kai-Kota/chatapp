package main

import (
	"chatapp/backend/infra"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	// ルートやハンドラの設定をここに追加
	return router
}

func main() {
	//インフラの構築
	infra.Initialize()
	db := infra.SetupDB()
	router := setupRouter(db)
	fmt.Println("Server is running on port 8080")

	router.Run(":8080")
}
