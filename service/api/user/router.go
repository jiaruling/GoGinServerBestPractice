package user

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/service/api/user/controller"
)

func Router() {
	global.GinRouter.V1.Any("/stu/*id", controller.Stus)
}
