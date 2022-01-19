package main

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/initial"
	"GoGinServerBestPractice/service/api"
	"GoGinServerBestPractice/service/backend_task"
	"GoGinServerBestPractice/utils"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

/*
   功能说明: 入口函数
   参考:
   创建人: 贾汝凌
   创建时间: 2022/01/18 10:29
*/

func main() {
	var (
		err  error
		quit chan os.Signal
	)
	// 初始化日志
	initial.InitCreateDir()
	initial.InitLog(global.LogPath)
	log.Println("1. 初始化日志成功")
	log.Println("2. 初始化线程数, 线程数量和cpu核数相等")
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("CPUNUM: ", runtime.NumCPU())
	log.Println("GOOS: ", runtime.GOOS)
	log.Println("3. 加载配置文件")
	if err = utils.ParseConfig("./config/server.yaml", &global.Config); err != nil {
		log.Fatalln("加载配置文件失败: ", err.Error())
	}
	log.Println("4. 连接数据库")
	initial.InitDB()
	log.Println("5. 多路复用器添加全局中间件")
	initial.InitGlobalMiddleware()
	log.Println("6. 注册路由")
	api.RegisterRouter()
	//log.Println(global.Config)
	log.Println("7. 启动http服务")
	initial.InitService()
	log.Println("8. 启动后台定时任务")
	backend_task.InitCelery()
	log.Println("-------------------------------------------------------------------------------------------------")
	// 优雅退出
	quit = make(chan os.Signal) // 定义一个无缓冲的通道
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//从quit中接收值，忽略结果
	<-quit
	log.Println("-------------------------------------------------------------------------------------------------")
	log.Println("优雅退出...")
	log.Println("资源重置, 保存数据...")
	log.Println("注销服务...")
	log.Println("程序优雅退出...")
	return
}
