package base

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/service/api/base/controller"
	"net/http"
)

func Router() {
	// 配置静态文件
	global.GinRouter.Router.StaticFS("/static", http.Dir(global.FILEPATH))
	global.GinRouter.Router.StaticFS("/log", http.Dir(global.LogFilePath))
	global.GinRouter.Router.NoRoute(controller.NotFound)      // 404 路由
	global.GinRouter.Router.GET("/health", controller.Health) // 服务健康检查
}
