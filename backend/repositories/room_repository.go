package repositories

import (
	"chatapp/backend/models"

	"gorm.io/gorm"
)

type IRoomRepository interface {
	Create(newRoom models.Room) (*models.Room, error)
	FindAll() (*[]models.Room, error)
}

type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) IRoomRepository {
	return &RoomRepository{db: db}
}

// チャットルームの作成
func (r *RoomRepository) Create(newRoom models.Room) (*models.Room, error) {
	//作成時にエラーが出たらnilとエラーを返す
	if err := r.db.Create(&newRoom).Error; err != nil {
		return nil, err
	}
	//成功したら作成したチャットルームを返す
	return &newRoom, nil
}

func (r *RoomRepository) FindAll() (*[]models.Room, error) {
	var rooms []models.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return &rooms, nil
}
