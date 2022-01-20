package grf

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 生成sql
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 16:24
*/

// *********************************************************************************************************************
// 新增记录
func GenInsertSQL(m interface{}, table string, CreatedFields, CreatedIgnoreFields, CreatedSetTimeFields []string) (sql string) {
	now := time.Now().Unix()
	field := ""
	value := ""
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	if len(CreatedFields) == 0 {
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := t.Elem().Field(i).Tag.Get("db")
			val := v.Elem().Field(i).Interface()
			if tag != "" {
				// 忽略字段
				if in(tag, CreatedIgnoreFields) {
					continue
				}
				// 设置时间
				if in(tag, CreatedSetTimeFields) {
					val = now
				}
				// 是否是第一个字段
				if field == "" {
					field = tag
					value = fmt.Sprintf("'%v'", val)
				} else {
					field += fmt.Sprintf(", %s", tag)
					value += fmt.Sprintf(", '%v'", val)
				}
			}
		}
	} else {
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := t.Elem().Field(i).Tag.Get("db")
			val := v.Elem().Field(i).Interface()
			if tag != "" {
				// 忽略字段
				if !in(tag, CreatedFields) {
					continue
				}
				// 设置时间
				if in(tag, CreatedSetTimeFields) {
					val = now
				}
				// 是否是第一个字段
				if field == "" {
					field = tag
					value = fmt.Sprintf("'%v'", val)
				} else {
					field += fmt.Sprintf(", %s", tag)
					value += fmt.Sprintf(", '%v'", val)
				}
			}
		}
	}
	return "INSERT INTO " + table + "(" + field + ")" + " VALUES" + "(" + value + ");"
}

// ******************************************************************************************************************END

// *********************************************************************************************************************

// 物理删除记录
func GenDeleteSQL(m interface{}, table string, id int64) (sql string) {
	return "DELETE FROM " + table + " WHERE id=" + fmt.Sprintf("%v;", id)
}

// 逻辑删除记录
func GenSoftDeleteSQL(m interface{}, table string, id int64, deleteField string) (sql string) {
	now := time.Now().Unix()
	t := reflect.TypeOf(m)
	for i := 0; i < t.Elem().NumField(); i++ {
		tag := t.Elem().Field(i).Tag.Get("db")
		if tag == deleteField {
			return "UPDATE " + table + fmt.Sprintf(" SET %s=%d ", deleteField, now) +
				fmt.Sprintf("WHERE id=%v AND %s is null;", id, deleteField)
		}
	}
	return
}

// ******************************************************************************************************************END

// *********************************************************************************************************************

// 更新记录
func GenUpdateSQL(m interface{}, table string, id int64, UpdateFields, UpdateIgnoreFields, UpdateSetTimeFields []string, deleteField string) (sql string) {
	now := time.Now().Unix()
	fieldAndValue := ""
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)
	if len(UpdateFields) == 0 {
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := t.Elem().Field(i).Tag.Get("db")
			val := v.Elem().Field(i).Interface()
			if tag != "" {
				if in(tag, UpdateIgnoreFields) {
					continue
				}
				if in(tag, UpdateSetTimeFields) {
					val = now
				}
				if fieldAndValue == "" {
					fieldAndValue = fmt.Sprintf("%s='%v'", tag, val)
				} else {
					fieldAndValue += fmt.Sprintf(", %s='%v'", tag, val)
				}
			}
		}
	} else {
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := t.Elem().Field(i).Tag.Get("db")
			val := v.Elem().Field(i).Interface()
			if tag != "" {
				if !in(tag, UpdateFields) {
					continue
				}
				if in(tag, UpdateSetTimeFields) {
					val = now
				}
				if fieldAndValue == "" {
					fieldAndValue = fmt.Sprintf("%s='%v'", tag, val)
				} else {
					fieldAndValue += fmt.Sprintf(", %s='%v'", tag, val)
				}
			}
		}
	}
	if deleteField == "" {
		return "UPDATE " + table + " SET " + fieldAndValue + fmt.Sprintf(" where id = %d;", id)
	}
	return "UPDATE " + table + " SET " + fieldAndValue + fmt.Sprintf(" where id = %d and %s is null;", id, deleteField)
}

// ******************************************************************************************************************END

// *********************************************************************************************************************
// 查询多条记录
func GenGetListSQL(m interface{}, table string, page, pageSize int64, condition, order string, fields, exFields []string, deleteField, all string) (sql string) {
	field := QueryFields(m, fields, exFields)
	if deleteField == "" || all != "" {
		if condition == "" {
			return "SELECT " + field + " FROM " + table + fmt.Sprintf(" %s LIMIT %d OFFSET %d;", order, pageSize, (page-1)*pageSize)
		}
		return "SELECT " + field + " FROM " + table + fmt.Sprintf(" WHERE %s %s LIMIT %d OFFSET %d;", strings.Trim(condition, " AND"), order, pageSize, (page-1)*pageSize)
	}
	return "SELECT " + field + " FROM " + table + fmt.Sprintf(" WHERE %s is null %s %s LIMIT %d OFFSET %d;", deleteField, condition, order, pageSize, (page-1)*pageSize)
}

// 根据ID查询一条记录
func GenGetByIdSQL(m interface{}, table string, id int64, fields, exFields []string, deleteField, all string) (sql string) {
	field := QueryFields(m, fields, exFields)
	if deleteField == "" || all != "" {
		return "SELECT " + field + " FROM " + table + fmt.Sprintf(" WHERE id=%d;", id)
	}
	return "SELECT " + field + " FROM " + table + fmt.Sprintf(" WHERE id=%d and %s is null;", id, deleteField)
}

// 查询记录总数
func GenGetTotalSQL(table string, condition, deleteField, all string) (sql string) {
	if deleteField == "" || all != "" {
		if condition == "" {
			return "SELECT count(1) as total FROM " + table + ";"
		}
		return "SELECT count(1) as total FROM " + table + " WHERE " + strings.Trim(condition, " AND") + ";"
	}
	return "SELECT count(1) as total FROM " + table + " WHERE " + deleteField + " is null" + condition + ";"
}

// ******************************************************************************************************************END

// *********************************************************************************************************************
// 查询字段
func QueryFields(m interface{}, fields, exFields []string) (field string) {
	t := reflect.TypeOf(m)
	if len(fields) == 0 {
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := strings.Split(t.Elem().Field(i).Tag.Get("json"), ",")[0]
			re := t.Elem().Field(i).Tag.Get("binding")
			ty := t.Elem().Field(i).Type.Name()
			if inExFields(tag, exFields) {
				continue
			}
			field = handlerField(tag, re, ty, field)
		}
	} else {
		for i := 0; i < t.Elem().NumField(); i++ {
			tag := strings.Split(t.Elem().Field(i).Tag.Get("json"), ",")[0]
			re := t.Elem().Field(i).Tag.Get("binding")
			ty := t.Elem().Field(i).Type.Name()
			if inFields(tag, fields) {
				field = handlerField(tag, re, ty, field)
			}
		}
	}
	return
}

func handlerField(tag, re, ty, field string) string {
	if in("required", strings.Split(re, ",")) {
		if tag != "" {
			if field == "" {
				field = tag
			} else {
				field += ", " + tag
			}
		}
	} else {
		switch ty {
		case "string", "Time":
			if tag != "" {
				if field == "" {
					field = fmt.Sprintf("ifnull(%s, \"\") as %s", tag, tag)
				} else {
					field += ", " + fmt.Sprintf("ifnull(%s, \"\") as %s", tag, tag)
				}
			}
		default:
			if tag != "" {
				if field == "" {
					field = fmt.Sprintf("ifnull(%s, 0) as %s", tag, tag)
				} else {
					field += ", " + fmt.Sprintf("ifnull(%s, 0) as %s", tag, tag)
				}
			}
		}
	}
	return field
}

// 查询条件
func ConditionSQL(c *gin.Context, search interface{}) (condition string) {
	t := reflect.TypeOf(search)
	k := t.Kind()
	switch k {
	case reflect.Slice:
		for _, item := range search.([]string) {
			if q := c.DefaultQuery(item, ""); q != "" {
				condition += fmt.Sprintf(" AND %s='%s'", item, q)
			}
		}
	default:
	}
	return
}

// 排序
func OrderSQL(c *gin.Context, order interface{}) (o string) {
	if orders := c.DefaultQuery("order", ""); orders != "" {
		orderList := strings.Split(orders, ",")
		for _, item := range orderList {
			if o == "" {
				o = "ORDER BY "
			}
			if item[0:1] == "-" {
				item = strings.Trim(item, "-")
				if o == "ORDER BY " {
					o += item + " DESC"
				} else {
					o += ", " + item + " DESC"
				}
			} else {
				if o == "ORDER BY " {
					o += item + " ASC"
				} else {
					o += ", " + item + " ASC"
				}
			}
		}
	} else {
		t := reflect.TypeOf(order)
		k := t.Kind()
		switch k {
		case reflect.Slice:
			for _, item := range order.([]string) {
				if o == "" {
					o = "ORDER BY "
				}
				if item[0:1] == "-" {
					item = strings.Trim(item, "-")
					if o == "ORDER BY " {
						o += item + " DESC"
					} else {
						o += ", " + item + " DESC"
					}
				} else {
					if o == "ORDER BY " {
						o += item + " ASC"
					} else {
						o += ", " + item + " ASC"
					}
				}
			}
		}
	}
	return
}
