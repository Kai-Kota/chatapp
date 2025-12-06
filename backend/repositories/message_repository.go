package repositories

import (
	"chatapp/backend/models"
	"fmt"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	CreateMessage(message models.Message) (*models.Message, error)
	FindAllMessages(roomID uint) (*[]models.Message, error)
	AssociateRoomToMessage(roomId uint, message *models.Message) error
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(message models.Message) (*models.Message, error) {
	if err := r.db.Create(&message).Error; err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessageRepository) FindAllMessages(roomID uint) (*[]models.Message, error) {
	var messages []models.Message
	if err := r.db.Joins("JOIN room_messages ON room_messages.message_id = messages.id").
		Where("room_messages.room_id = ?", roomID).
		Find(&messages).Error; err != nil {
		return nil, err
	}
	return &messages, nil
}

func (r *MessageRepository) AssociateRoomToMessage(roomId uint, message *models.Message) error {
	if err := r.db.Model(&models.Room{Model: gorm.Model{ID: roomId}}).Association("Messages").Append(message); err != nil {
		fmt.Println(roomId, message.ID)
		return err
	}
	return nil
}
