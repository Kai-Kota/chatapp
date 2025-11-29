package repositories

import (
	"chatapp/backend/models"

	"gorm.io/gorm"
)

type IRoomRepository interface {
	Create(newRoom models.Room) (*models.Room, error)
	GetUserRooms(userId uint) (*[]models.Room, error)
	FindUserIdByName(userName string) uint
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

func (r *RoomRepository) GetUserRooms(userId uint) (*[]models.Room, error) {
	var rooms []models.Room
	if err := r.db.Where("member1 = ? OR member2 = ?", userId, userId).Find(&rooms).Error; err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (r *RoomRepository) FindUserIdByName(userName string) uint {
	var user models.User
	r.db.First(&user, "user_name = ?", userName)
	return user.ID
}
