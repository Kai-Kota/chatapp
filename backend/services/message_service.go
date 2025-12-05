package services

import (
	"chatapp/backend/models"
	"chatapp/backend/repositories"
	"fmt"
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

	message, err := s.repository.CreateMessage(newMessage)
	if err != nil {
		return nil, err
	}

	fmt.Println(message.ID)
	//メッセージとルームの関連付け
	if err := s.repository.AssociateRoomToMessage(roomId, message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *MessageService) GetRoomMessages(roomId uint) (*[]models.Message, error) {
	return s.repository.FindAllMessages(roomId)
}
