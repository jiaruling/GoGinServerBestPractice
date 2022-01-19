package initial

import (
	"GoHttpServerBestPractice/global"
	middleware2 "GoHttpServerBestPractice/service/middleware"
)

/*
   功能说明: 添加全局中间件
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 15:42
*/

func InitGlobalMiddleware() {
	middleware := []global.Middleware{
		middleware2.Cors,
		middleware2.AccessLog,
	}
	global.Mux.Use(middleware...)
}
