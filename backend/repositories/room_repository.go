package repositories

import (
	"chatapp/backend/models"
	"fmt"

	"gorm.io/gorm"
)

type IRoomRepository interface {
	Create(newRoom models.Room) (*models.Room, error)
	GetUserRooms(userId uint) (*[]models.Room, error)
	AssoiciateUserToRoom(userId uint, room *models.Room) error
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

	// JOIN を使って指定ユーザーが所属する rooms を取得する
	if err := r.db.Joins("JOIN user_rooms ON user_rooms.room_id = rooms.id").
		Where("user_rooms.user_id = ?", userId).
		Find(&rooms).Error; err != nil {
		return nil, err
	}

	return &rooms, nil
}

func (r *RoomRepository) FindUserIdByName(userName string) uint {
	var user models.User
	r.db.First(&user, "user_name = ?", userName)
	return user.ID
}

func (r *RoomRepository) AssoiciateUserToRoom(userId uint, room *models.Room) error {
	//チャットルームの作成後、ユーザーとチャットルームの関連付けを行う
	fmt.Println(room)
	if err := r.db.Model(&models.User{Model: gorm.Model{ID: userId}}).Association("Rooms").Append(room); err != nil {
		return err
	}
	return nil
}
