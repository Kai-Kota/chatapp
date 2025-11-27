package dto

type SignupInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}