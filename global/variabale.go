package global

import (
	. "GoGinServerBestPractice/global/config_struct"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
)

/*
   功能说明: 变量
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 10:52
*/

type Router struct {
	Router *gin.Engine
	V1     *gin.RouterGroup
}

var (
	Config      ServerConfig
	GinRouter   *Router
	Trans       ut.Translator
	RDB         *gorm.DB
	WDB         *gorm.DB
	Validate    *validator.Validate
	Expires     time.Duration
	ETicker     *time.Ticker
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
	StartTime   string
	Version     string
)

// 初始化全局变量
func init() {
	Config.Log = Log{
		LogDir:     "./log",
		InfoLog:    "info.log",
		ErrorLog:   "error.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   false,
	}
	Validate = validator.New()
	Expires = 10 // 10s
	ETicker = time.NewTicker(Expires * time.Second)
	StartTime = time.Now().Format("2006.01.02 15:04:05")
	Version = os.Getenv(Config.SystemVersion)
}
