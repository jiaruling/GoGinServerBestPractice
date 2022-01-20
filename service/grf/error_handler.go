package grf

import (
	"GoGinServerBestPractice/global/errInfo"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 错误处理
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:56
*/

func Handler200(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": errInfo.SUCCESS, "data": data})
	return
}

func Handler201(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "msg": errInfo.SUCCESS, "data": data})
	return
}

func Handler204(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusNoContent, "msg": errInfo.SUCCESS, "data": nil})
	return
}

func Handler400(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": msg, "data": data})
	return
}

func Handler404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "msg": errInfo.NotFound, "data": nil})
	return
}

func Handler500(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "msg": msg, "data": data})
	return
}
