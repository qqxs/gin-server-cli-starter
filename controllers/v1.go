package controllers

import (
	"gin-server-cli/services"
	"gin-server-cli/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostLogin 登录控制器
func PostLogin(c *gin.Context) {
	service := services.UserLoginService{}

	if err := c.ShouldBind(&service); err == nil {
		res := service.UserLogin()
		c.JSON(http.StatusOK, res)
	} else {
		// err.Error()
		c.JSON(http.StatusOK, utils.BuildBadResponse("无效参数"))
	}
}

// PostRegister 注册成为用户接口
func PostRegister(c *gin.Context) {
	service := services.UserCreateService{}

	if err := c.ShouldBind(&service); err == nil {
		res := service.UserCreate()
		c.JSON(http.StatusOK, res)
	} else {
		log.Println(err.Error())
		c.JSON(http.StatusOK, utils.BuildBadResponse("注册失败"))
	}
}
