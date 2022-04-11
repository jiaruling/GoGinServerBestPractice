package initial

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/service/grf"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

/*
   功能说明: 初始化数据库
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 9:56
*/

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 15:47
*/

func InitDB() {
	MySQL := global.Config.MySQL
	DNS := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?%s", MySQL.User, MySQL.Password, MySQL.Ip, MySQL.Port, MySQL.Db, MySQL.Parameter)
	database, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		global.SugarLogger.Error("open mysql failed,", err.Error())
		os.Exit(1)
	}
	// 设置数据库连接池
	sqlDB, err := database.DB()
	if err != nil {
		global.SugarLogger.Error("设置数据库连接池失败,", err.Error())
		os.Exit(1)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	//defer database.Close()  // 注意这行代码要写在上面err判断的下面
	grf.RDB = database
	grf.WDB = database
	grf.GlobalPageMax = 5
	grf.GlobalPageMin = 1
	global.RDB = database
	global.WDB = database
}
