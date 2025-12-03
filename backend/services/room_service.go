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

	// 新しいチャットルームの初期データを設定
	newRoom := models.Room{
		Name: "Room_" + "_" + pertner,
		//Messages: []models.Message{}, // 初期状態ではメッセージは空
	}

	partnerId := s.repository.FindUserIdByName(pertner)

	// チャットルームの作成
	room, err := s.repository.Create(newRoom)

	// ユーザーとチャットルームの関連付け
	if err := s.repository.AssoiciateUserToRoom(userId, room); err != nil {
		return nil, err
	}
	if err := s.repository.AssoiciateUserToRoom(partnerId, room); err != nil {
		return nil, err
	}

	return room, err
}

func (s *RoomService) GetUserRooms(userId uint) (*[]models.Room, error) {
	return s.repository.GetUserRooms(userId)
}
