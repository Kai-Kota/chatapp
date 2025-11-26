package services

import (
	"chatapp/backend/dto"
	"chatapp/backend/models"
	"chatapp/backend/repositories"
)

type IRoomService interface {
	Create(createRoomInput dto.CreateRoomInput) (*models.Room, error)
}

type RoomService struct {
	repository repositories.IRoomRepository
}

func NewRoomService(repository repositories.IRoomRepository) IRoomService {
	return &RoomService{repository: repository}
}

func (s *RoomService) Create(createRoomInput dto.CreateRoomInput) (*models.Room, error) {
	// 新しいチャットルームの初期データを設定
	newRoom := models.Room{
		RoomID: createRoomInput.RoomID,  
		//Messages: []models.Message{}, // 初期状態ではメッセージは空
	}
	return s.repository.Create(newRoom)
}
