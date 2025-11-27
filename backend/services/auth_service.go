package services

import (
	"chatapp/backend/models"
	"chatapp/backend/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(userName string, password string) error
	Login(email string, password string) (*string, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(userName string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		UserName: userName,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

func (s *AuthService) Login(userName string, password string) (*string, error) {
	
	//ユーザーが存在するか確認
	founduser, err := s.repository.FindUser(userName)
	if err != nil {
		return nil, err
	}

	//パスワードの確認
	err = bcrypt.CompareHashAndPassword([]byte(founduser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	//トークンの作成
	token, err := CreateToken(founduser.ID, founduser.UserName)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(userId uint, userName string) (*string, error) {
	// トークンのペイロードを設定
	claims := jwt.MapClaims{
		"sub":      userId,
		"userName": userName,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}
	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名して文字列を取得
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
