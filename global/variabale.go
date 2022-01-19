package global

import (
	. "GoGinServerBestPractice/global/config_struct"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

/*
   功能说明: 变量
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 10:52
*/

var (
	Config    ServerConfig
	RDB       *gorm.DB
	WDB       *gorm.DB
	AccessLog *log.Logger
	SqlLog    *log.Logger
	TaskLog   *log.Logger
	LogPath   []string
	Server    *http.Server
	Validate  *validator.Validate
	Expires   time.Duration
	ETicker   *time.Ticker
)

// 初始化全局变量
func init() {
	LogPath = []string{"./log/logs.log", "./log/access.log", "./log/sql.log", "./log/backend_task.log"}
	Validate = validator.New()
	Expires = 10 // 10s
	ETicker = time.NewTicker(Expires * time.Second)
}
