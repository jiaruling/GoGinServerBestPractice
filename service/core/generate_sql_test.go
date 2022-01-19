package core

import (
	"fmt"
	"strings"
	"testing"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/15 12:23
*/

type Student struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	Age         int64  `json:"age" db:"age" binding:"required"`
	Class       string `json:"class" db:"class" binding:"required"`
	CreatedAt   int64  `json:"created_at" db:"created_at"`
	UpdatedAt   int64  `json:"updated_at" db:"updated_at"`
	DeletedAtAt int64  `json:"deleted_at" db:"deleted_at"`
}

func TestGenerateSql(t *testing.T) {
	//database-sql := GenInsertSQL(&Student{Name: "张三", Age: 18, Class:"大二"}, "student")
	//log.Println(database-sql)
	//database-sql = GenUpdateSQL(&Student{Name: "张三", Age: 18, Class:"大二"}, "student", 2)
	//log.Println(database-sql)
	//database-sql = GenGetByIdSQL(&Student{Name: "张三", Age: 18, Class:"大二"}, "student", 2)
	//log.Println(database-sql)
	//database-sql = GenGetListSQL(&Student{Name: "张三", Age: 18, Class:"大二"}, "student", 1, 10)
	//log.Println(database-sql)
	//ConditionSQL([]string{"name", "age"})
	//s := "hello yes"
	//fmt.Println(strings.Trim(s, "he"))
	//fmt.Println(strings.Trim(s, "es"))
	//fmt.Println(strings.Trim(s, "ll"))
	s := "id,omitempty"
	fmt.Println(strings.Split(s, ",")[0])
}