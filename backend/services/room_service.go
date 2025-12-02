package services

import (
	"chatapp/backend/models"
	"chatapp/backend/repositories"
)

type IRoomService interface {
	Create(userId uint, pertner string) (*models.Room, error)
	GetUserRooms(userId uint) (*[]models.Room, error)
}

type RoomService struct {
	repository repositories.IRoomRepository
}

func NewRoomService(repository repositories.IRoomRepository) IRoomService {
	return &RoomService{repository: repository}
}

func (s *RoomService) Create(userId uint, pertner string) (*models.Room, error) {
	pertnerId := s.repository.FindUserIdByName(pertner)

	// 新しいチャットルームの初期データを設定
	newRoom := models.Room{
		
		//Messages: []models.Message{}, // 初期状態ではメッセージは空
	}
	return s.repository.Create(newRoom)
}

func (s *RoomService) GetUserRooms(userId uint) (*[]models.Room, error) {
	return s.repository.GetUserRooms(userId)
}
