package core

import (
	"fmt"
	"log"
	"os"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2021/12/15 13:41
*/
func init() {
	logFile, err := os.OpenFile(LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	//defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}