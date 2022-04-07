package main

import (
	"net/http"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	zap+日志分级分文件+按大小切割日志
*/

var Logger *zap.Logger

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	simpleHttpGet("www.topgoer.com")
	simpleHttpGet("http://www.topgoer.com")
}

func InitLogger() {
	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	encoder := getEncoder()

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter("./demo.log")
	warnWriter := getWriter("./demo_error.log")

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	for i := 0; i < 100000; i++ {
		sugarLogger.Debugf("Trying to hit GET request for %s", url)
		resp, err := http.Get(url)
		if err != nil {
			sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
		} else {
			sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
			resp.Body.Close()
		}
	}
}
