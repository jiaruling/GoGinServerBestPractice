package initial

import (
	"GoGinServerBestPractice/global"
	"os"
	"strconv"
)

/*
   功能说明: 启动服务
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 14:45
*/

func InitService() {
	go func() {
		if err := global.GinRouter.Router.Run(global.Config.Service.Addr + ":" + strconv.Itoa(global.Config.Service.Port)); err != nil {
			global.SugarLogger.Error("<main>: http服务启动失败: ", err.Error())
			os.Exit(1)
		}
	}()
}
