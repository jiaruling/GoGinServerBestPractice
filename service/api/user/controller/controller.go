package controller

import (
	"GoGinServerBestPractice/service/api/user/dao"
	"GoGinServerBestPractice/service/grf"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/19 16:46
*/

func Stus(c *gin.Context) {
	// 通过全局变量赋值给局部变量，可以实现并发
	s := dao.Stu
	s.M = new(dao.Student)
	grf.Dispatcher(c, s)
}
