package initial

import (
	"GoGinServerBestPractice/global"
	"GoGinServerBestPractice/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	logList, err := WalkDir(global.Config.Log.LogDir, "log")
	if err != nil {
		fmt.Println("获取日志文件名称失败")
	} else {
		deleteLogFile(logList)
	}

	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	encoder := getEncoder()

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(global.Config.Log.LogDir + "/" + global.Config.Log.InfoLog)
	warnWriter := getWriter(global.Config.Log.LogDir + "/" + global.Config.Log.ErrorLog)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	global.SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    global.Config.Log.MaxSize, //日志文件的最大大小（以MB为单位）
		MaxBackups: global.Config.Log.MaxBackups,
		MaxAge:     global.Config.Log.MaxAge,
		Compress:   global.Config.Log.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 删除文件
func deleteLogFile(LogList []string) {
	for _, item := range LogList {
		if utils.Exists(item) {
			_ = os.Remove(item)
		}
	}
}

// 获取文件夹下指定文件
func WalkDir(dir, suffix string) (files []string, err error) {
	files = []string{}
	err = filepath.Walk(dir, func(fname string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			//忽略目录
			return nil
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			//文件后缀匹配
			files = append(files, fname)
		}
		return nil
	})
	return files, err
}
