package controller

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/service/core"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/19 16:46
*/

func NotFound(c *gin.Context) {
	core.Handler404(c)
	return
}

func Health(c *gin.Context) {
	core.Handler200(c, global.Config.Name)
	return
}
