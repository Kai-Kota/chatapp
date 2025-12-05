package services

import (
	"chatapp/backend/models"
	"chatapp/backend/repositories"
)

type IMessageService interface {
	CreateMessage(userId uint, content string, roomId uint) (*models.Message, error)
	GetRoomMessages(roomId uint) (*[]models.Message, error)
}

type MessageService struct {
	repository repositories.IMessageRepository
}

func NewMessageService(repository repositories.IMessageRepository) IMessageService {
	return &MessageService{repository: repository}
}

func (s *MessageService) CreateMessage(userId uint, content string, roomId uint) (*models.Message, error) {

	newMessage := models.Message{
		UserID:  userId,
		Content: content,
	}
	//メッセージの作成
	if err := s.repository.CreateMessage(newMessage); err != nil {
		return nil, err
	}

	//メッセージとルームの関連付け
	if err := s.repository.AssociateRoomToMessage(roomId, &newMessage); err != nil {
		return nil, err
	}

	return &newMessage, nil
}

func (s *MessageService) GetRoomMessages(roomId uint) (*[]models.Message, error) {
	return s.repository.FindAllMessages(roomId)
}