package main

import (
	"chatapp/backend/infra"
	"chatapp/backend/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// マイグレーションの実行(データベースに登録)
	if err := db.AutoMigrate(
		&models.User{}, &models.Room{}, &models.Message{}); err != nil {
		panic("failed to migrate database")
	}
}