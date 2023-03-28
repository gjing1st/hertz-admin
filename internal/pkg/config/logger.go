package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitLogger(level string, output string, dir string, caller bool) {
	switch level {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	format := log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	}
	log.SetFormatter(&format)

	switch output {
	case "std":
		writer := os.Stdout
		log.SetOutput(writer)
	case "file":
		//writer, _ := rotatelogs.New(
		//	dir+"/log/mck.log",
		//	//rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		//	rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		//	rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
		//	rotatelogs.WithRotationCount(3),            //设置3份 大于3份 或到了清理时间 开始清理
		//	//rotatelogs.WithRotationSize(100*1024*1024), //设置100MB大小,当大于这个容量时，创建新的日志文件
		//)
		//
		//log.SetOutput(writer)

		logger := &lumberjack.Logger{
			LocalTime:  true,
			Filename:   dir + "/mck.log",
			MaxSize:    200, // 一个文件最大为200M
			MaxBackups: 5,   // 最多同时保存5份文件
			MaxAge:     7,   // 一个文件最多同时存在7天
			Compress:   false,
		}
		log.SetOutput(logger)
	}

	log.SetReportCaller(caller)
}

func LogError(fields log.Fields, args ...interface{}) {
	log.WithFields(fields).Error(args...)
}

func WriteDataLogs(info, err interface{}) log.Fields {
	return log.Fields{
		"event_info": info,
		"err":        err,
	}
}
