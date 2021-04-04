package router

import (
	"gin-server-cli/controllers"

	"github.com/gin-gonic/gin"
)

func registerWeb(app *gin.Engine) {

	// 静态资源地址
	// https://github.com/gin-gonic/gin/blob/master/routergroup.go#L166
	app.Static("/static", "./public/static")

	// gin set func map
	// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/29.html
	// https://github.com/gin-gonic/gin/blob/master/gin.go#L220
	// app.SetFuncMap(template.FuncMap{
	// 	"is_debug":  utils.IsDebug,   // is debug -> bool
	// 	"unescaped": utils.Unescaped, // string -> html
	// })

	// all page file path
	// https://github.com/gin-gonic/gin/blob/master/gin.go#L183
	app.LoadHTMLGlob("./public/*.html")

	// 首页 router /
	app.GET("/", controllers.IndexWeb)
	app.GET("/index.html", controllers.IndexWeb)
	app.GET("/index", controllers.IndexWeb)
	// 登录 app /login.html
	app.GET("/login.html", controllers.LoginWeb)
	app.GET("/login", controllers.LoginWeb)
	// 注册 app /register.html
	app.GET("/register.html", controllers.RegisterWeb)
	app.GET("/register", controllers.RegisterWeb)

	// 404
	app.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	// 404
	app.NoMethod(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
}
