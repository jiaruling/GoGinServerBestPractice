package main

import (
	"net/http"

	"go.uber.org/zap"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/4/7 13:57
*/
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.5lmh.com")
	simpleHttpGet("http://www.google.com")
}

func InitLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
