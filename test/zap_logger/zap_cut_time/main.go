package main

import (
	"errors"
	"strconv"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	zap+日志分级分文件+按时间切割日志
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
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	// 日志文件大于1M分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithRotationSize(1024*1024*1), // 1M
	)

	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(hook)
}

func simpleHttpGet(url string) {
	for i := 0; i < 10000; i++ {
		sugarLogger.Debugf("Trying to hit GET request for %s", url)
		//resp, err := http.Get(url)
		if i%10 == 0 {
			sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, errors.New(strconv.Itoa(i)))
		} else {
			sugarLogger.Infof("Success! statusCode = %s for URL %s", "200 ok", url)
			//resp.Body.Close()
		}
	}
}
