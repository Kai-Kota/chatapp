package controllers

import (
	"chatapp/backend/dto"
	"chatapp/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Signup(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Signup(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.Signup(input.UserName, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	//JSONのバインド
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ログイン処理
	token, err := c.service.Login(input.UserName, input.Password)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", *token, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
