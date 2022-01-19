package initial

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/service/core"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
	database, err := sqlx.Open("mysql", DNS)
	if err != nil {
		log.Fatalln("open mysql failed,", err)
		return
	}
	//defer database.Close()  // 注意这行代码要写在上面err判断的下面
	core.RDB = database
	core.WDB = database
	core.GlobalPageMax = 5
	core.GlobalPageMin = 1
	global.RDB = database
	global.WDB = database
}
