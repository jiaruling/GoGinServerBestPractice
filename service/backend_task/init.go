package backend_task

/*
   功能说明: 后台常驻定时任务
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 11:52
*/

func InitCelery() {
	go task1()
	//go task2()
	//go task3()
}
