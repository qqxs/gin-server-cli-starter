package services

import (
	"gin-server-cli/utils"

	"github.com/gin-gonic/gin"
)

// LoginUserService bind login data
type UserLoginService struct {
	ID       uint   `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"required,email,min=4,max=200"`
	Password string `json:"password" form:"password" binding:"required,min=5,max=100"`
	// CaptchaID string `json:"captchaId" form:"captchaId" binding:"required"`
	// Captcha   string `json:"captcha" form:"captcha" binding:"required"`
}

// CreateUserService bind register data
type UserCreateService struct {
	UserLoginService
	Name   string `json:"name" form:"name" binding:"required,min=1,max=100"`
	Avater string `json:"avatar" form:"avatar"`
}

// UserLoginService UserLogin 用户登录
func (service *UserLoginService) UserLogin() utils.ResponseResult {
	return utils.BuildOKResponse(gin.H{
		"token": "this is token",
	})
}

// UserCreateService UserCreate 创建用户
func (service *UserCreateService) UserCreate() utils.ResponseResult {
	return utils.BuildOKResponse(true)
}
