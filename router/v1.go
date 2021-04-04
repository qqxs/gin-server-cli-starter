package router

import (
	"gin-server-cli/controllers"

	"github.com/gin-gonic/gin"
)

func registerRouterV1(app *gin.RouterGroup) {
	// curl -d "email=xiaoshaoqq@gmail.com&password=1234567" http://localhost:8080/api/v1/login
	app.POST("/login", controllers.PostLogin)

	// curl -d "email=xiaoshaoqq@gmail.com&password=1234567&name=ShineShao" http://localhost:8080/api/v1/register
	app.POST("/register", controllers.PostRegister)
}
