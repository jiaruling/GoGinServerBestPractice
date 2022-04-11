package main

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/initial"
	"GoGinServerBestPractice/service/api"
	"GoGinServerBestPractice/service/backend_task"
	"GoGinServerBestPractice/utils"
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
	initial.InitLogger()
	global.SugarLogger.Info("<main>: ", "1. 初始化日志成功")
	global.SugarLogger.Info("<main>: ", "2. 初始化线程数, 线程数量和cpu核数相等")
	runtime.GOMAXPROCS(runtime.NumCPU())
	global.SugarLogger.Info("<main>: ", "CPUNUM: ", runtime.NumCPU())
	global.SugarLogger.Info("<main>: ", "GOOS: ", runtime.GOOS)
	global.SugarLogger.Info("<main>: ", "3. 加载配置文件")
	if err = utils.ParseConfig("./config/server.yaml", &global.Config); err != nil {
		global.SugarLogger.Error("<main>: ", "加载配置文件失败: ", err.Error())
		os.Exit(1)
	}
	global.SugarLogger.Info("<main>: ", "4. 连接数据库")
	initial.InitDB()
	global.SugarLogger.Info("<main>: ", "5. 初始化验证器的翻译器")
	initial.InitTrans("zh")
	global.SugarLogger.Info("<main>: ", "6. 初始化自定义的验证器")
	if err = initial.InitValidator(); err != nil {
		global.SugarLogger.Error("<main>: ", "初始化自定义的验证器失败: ", err.Error())
	}
	global.SugarLogger.Info("<main>: ", "7. 初始化Gin")
	initial.InitGin()
	global.SugarLogger.Info("<main>: ", "8. 注册路由")
	api.RegisterRouter()
	//log.Println(global.Config)
	global.SugarLogger.Info("<main>: ", "9. 启动http服务")
	initial.InitService()
	global.SugarLogger.Info("<main>: ", "10. 启动后台定时任务")
	backend_task.InitBackendTask()
	// 优雅退出
	quit = make(chan os.Signal) // 定义一个无缓冲的通道
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//从quit中接收值，忽略结果
	<-quit
	global.SugarLogger.Info("<main>: ", "-1. 优雅退出...")
	global.SugarLogger.Info("<main>: ", "-2. 资源重置, 保存数据...")
	global.SugarLogger.Info("<main>: ", "-3. 注销服务...")
	global.SugarLogger.Info("<main>: ", "-4. 程序优雅退出...")
	return
}
