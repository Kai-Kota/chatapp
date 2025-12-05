package repositories

import (
	"chatapp/backend/models"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	CreateMessage(message models.Message) error
	FindAllMessages(roomID uint) (*[]models.Message, error)
	AssociateRoomToMessage(roomId uint, message *models.Message) error
}

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(message models.Message) error {
	if err := r.db.Create(&message).Error; err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) FindAllMessages(roomID uint) (*[]models.Message, error) {
	var messages []models.Message
	if err := r.db.Where("room_id = ?", roomID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return &messages, nil
}

func (r *MessageRepository) AssociateRoomToMessage(roomId uint, message *models.Message) error {
	if err := r.db.Model(&models.Message{Model: gorm.Model{ID: roomId}}).Association("Messages").Append(message); err != nil {
		return err
	}
	return nil
}
