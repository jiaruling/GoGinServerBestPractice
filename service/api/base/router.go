package base

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/service/api/base/controller"
	"net/http"
)

func Router() {
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	global.GinRouter.Router.StaticFS("/static", http.Dir(global.FILEPATH))
	global.GinRouter.Router.StaticFS("/log", http.Dir(global.LogFilePath))
	global.GinRouter.Router.NoRoute(controller.NotFound)      // 404 路由
	global.GinRouter.Router.GET("/health", controller.Health) // 服务健康检查
	global.GinRouter.Router.GET("/ping", controller.Ping)     // 服务健康检查
	global.GinRouter.Router.GET("/file/base64", controller.GetConfigFile)
	global.GinRouter.Router.POST("/file/base64", controller.PostConfigFile)
}
