package grf

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
   功能说明: 功能函数
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/15 18:04
*/

func InFields(field string, Fields []string) bool {
	return in(field, Fields)
}

func InExFields(field string, ExFields []string) bool {
	return in(field, ExFields)
}

// 判断元素是否在切片中
func in(f string, fl []string) bool {
	if len(fl) == 0 {
		return false
	}
	for _, item := range fl {
		if f == item {
			return true
		}
	}
	return false
}

// 分页器
func Paging(c *gin.Context, PageMax, PageMin int64) (page, pageSize int) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	// 页码
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}

	// 每一页大小
	pageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil {
		if PageMin > 0 {
			pageSize = int(PageMin)
		} else {
			pageSize = int(GlobalPageMin)
		}
	}
	if PageMax > 0 {
		if pageSize > int(PageMax) {
			pageSize = int(PageMax)
		}
	} else {
		if pageSize > int(GlobalPageMax) {
			pageSize = int(GlobalPageMax)
		}
	}
	if PageMin > 0 {
		if pageSize <= 0 {
			pageSize = int(PageMin)
		}
	} else {
		if pageSize <= 0 {
			pageSize = int(GlobalPageMin)
		}
	}
	return
}

/*
	去除字符串前面的结构体名称:
		{User.Name: "张三"} --> {Name: "张三"}
*/
func RemoveTopStruct(fields map[string]string) (b []byte, err error) {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	b, err = json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return b, nil
}
