package router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter init gin router
func InitRouter(app *gin.Engine) {

	// 设置日志文件
	// loggerFile()
	// 自定义日志配置
	// c := setLoggerConfig()

	// 中间件
	// router.Use(middleware.Cors())
	// router.Use(gin.LoggerWithConfig(c))
	// router.Use(gin.Recovery())

	// 注册web 页面
	registerWeb(app)

	// {
	// 	// 验证码 地址
	// 	router.GET("/captcha/:captchaID", controller.ServeHTTP)
	// }

	// // 注册api 接口
	api := app.Group("api")
	{
		registerRouterV1(api.Group("/v1"))
		// registerRouterV2(api.Group("/v2"))

		// auth := middlewares.JWTAuth()

		// registerAuthRouterV1(api.Group("/v1", auth))
		// registerAuthRouterV2(api.Group("/v2", auth))
	}

}

// // loggerFile set logger default writer
// func loggerFile() {
// 	_ = os.MkdirAll(conf.Config.GetString("logging.log_path")+conf.Config.GetString("logging.log_http_path"), os.ModePerm)
// 	_ = os.MkdirAll(conf.Config.GetString("logging.log_path")+conf.Config.GetString("logging.log_error_path"), os.ModePerm)
// 	now := time.Now()
// 	logFileName := now.Format("2006-01-02") + ".log"

// 	logFile, _ := os.OpenFile(conf.Config.GetString("logging.log_path")+conf.Config.GetString("logging.log_http_path")+logFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	errorLogFile, _ := os.OpenFile(conf.Config.GetString("logging.log_path")+conf.Config.GetString("logging.log_error_path")+logFileName, os.O_RDWR|os.O_CREATE, os.ModePerm)

// 	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
// 	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
// 	gin.DefaultErrorWriter = io.MultiWriter(errorLogFile, os.Stdout)
// }

// // setLoggerConfig set logger config
// func setLoggerConfig() gin.LoggerConfig {
// 	return gin.LoggerConfig{
// 		SkipPaths: []string{"/swagger/*"},
// 		Formatter: func(params gin.LogFormatterParams) string {
// 			return fmt.Sprintf("[%s] %s | %d | %12s | %15s | \"%s %s %s\" | %s | \" %s \" |\n",
// 				"GIN",
// 				params.TimeStamp.Format("2006/01/02 15:04:05"),
// 				params.StatusCode,
// 				params.Latency,
// 				params.ClientIP,
// 				params.Method,
// 				params.Path,
// 				params.Request.Proto,
// 				params.Request.UserAgent(),
// 				params.ErrorMessage,
// 			)
// 		},
// 	}
// }
