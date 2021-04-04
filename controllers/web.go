package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index web index.
// router [/]
func IndexWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin Demo Starter",
	})
}

// Login web login.
// router [/login.html]
func LoginWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录",
	})
}

// Register web register.
// router [/Register.html]
func RegisterWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "注册",
	})
}
