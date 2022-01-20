package controller

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/global/errInfo"
	"GoGinServerBestPractice/service/api/base/form"
	"GoGinServerBestPractice/service/grf"
	"GoGinServerBestPractice/utils"
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/19 16:46
*/

func NotFound(c *gin.Context) {
	grf.Handler404(c)
	return
}

func Health(c *gin.Context) {
	grf.Handler200(c, global.Config.Name)
	return
}

// 二进制文件转base64后传输
func GetConfigFile(c *gin.Context) {
	file := c.DefaultQuery("file", "")
	path := global.ConfigFilePath + file
	if !utils.Exists(path) {
		grf.Handler400(c, errInfo.FileNotFound, nil)
		return
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		grf.Handler500(c, errInfo.FileReadFailed, nil)
		return
	}
	grf.Handler200(c, base64.StdEncoding.EncodeToString(content))
	return
}

// base64数据解码后写入文件
func PostConfigFile(c *gin.Context) {
	var v form.FileContent
	if err := c.ShouldBind(&v); err != nil {
		grf.Handler400(c, errInfo.FormValidateFailed, nil)
		return
	}
	content, err := base64.StdEncoding.DecodeString(v.Content)
	if err != nil {
		grf.Handler500(c, errInfo.Base64DecodeFailed, nil)
		return
	}
	if err := ioutil.WriteFile(global.ConfigFilePath+v.FileName, content, os.ModePerm); err != nil {
		grf.Handler500(c, errInfo.FileWriteFailed, nil)
		return
	}
	grf.Handler200(c, nil)
	return
}
