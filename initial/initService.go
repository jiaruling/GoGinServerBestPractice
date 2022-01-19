package initial

import (
	"GoGinServerBestPractice/global"
	"log"
	"net/http"
)

/*
   功能说明: 启动服务
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 14:45
*/

func InitService() {
	go func() {
		if err := global.Server.ListenAndServe(); err != nil {
			log.Fatalln("http服务启动失败: ", err.Error())
		}
	}()
	go func() {
		if err := http.ListenAndServe(":8000", http.FileServer(http.Dir(global.FILEPATH))); err != nil {
			log.Fatalln("http静态资源服务启动失败: ", err.Error())
		}
	}()
	go func() {
		if err := http.ListenAndServe(":8001", http.FileServer(http.Dir(global.LogFilePath))); err != nil {
			log.Fatalln("http日志资源服务启动失败: ", err.Error())
		}
	}()
}
