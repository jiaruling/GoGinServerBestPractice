package grf

/*
   功能说明: 数据库操作
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/14 16:03
*/

// 增删改
func ExecDB(sql string) (lastId int64, err error) {
	result := WDB.Exec(sql)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// 列表查询
func getListDB(sql string) (list []map[string]interface{}, err error) {
	list = make([]map[string]interface{}, 0)
	result := RDB.Raw(sql).Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}

// 查询记录总数
func getTotalDB(sql string) (total int64) {
	_ = RDB.Raw(sql).Scan(&total)
	return
}

// 单条记录查询
func getByIdDB(sql string) (list []map[string]interface{}, err error) {
	list = make([]map[string]interface{}, 0)
	result := RDB.Raw(sql).Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
