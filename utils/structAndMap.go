package utils

import (
	"encoding/json"
	"reflect"
)

/*
   功能说明:
		- 结构体转map
		- 结构体列表转map列表
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/19 15:41
*/
func StructToMap(s interface{}, m *map[string]interface{}) (err error) {
	var (
		data []byte
	)
	data, err = json.Marshal(s) // s 为指针类型
	if err != nil {
		return
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		return
	}
	return
}

func StructListToMapList(s interface{}) (mList *[]*map[string]interface{}, err error) {
	var (
		data  []byte
		m     *map[string]interface{}
		MList []*map[string]interface{}
	)
	MList = make([]*map[string]interface{}, 0)
	// 判断类型
	if reflect.TypeOf(s).Kind() == reflect.Slice {
		obj := reflect.ValueOf(s)
		for i := 0; i < obj.Len(); i++ {
			m = new(map[string]interface{})
			ele := obj.Index(i)
			data, err = json.Marshal(ele.Interface())
			if err != nil {
				return
			}
			err = json.Unmarshal(data, m)
			if err != nil {
				return
			}
			MList = append(MList, m)
		}
	}
	return &MList, nil
}
