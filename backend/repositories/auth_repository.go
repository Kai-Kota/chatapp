package repositories

import (
	"chatapp/backend/models"
	"errors"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user models.User) error
	FindUser(userName string) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

// コンストラクタ
func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) error {
	//ユーザをデータベースに作成
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) FindUser(userName string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "user_name = ?", userName).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
