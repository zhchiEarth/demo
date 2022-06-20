package log

import (
	"compound/internal/conf"
	"compound/internal/pkg"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

func Logger(c *conf.Server) log.Logger {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	writer := getLogWriter(c.Log.File)

	// 写入文件的同时，写入控制台，供调试
	//if conf.EnableConsole {
	writer = io.MultiWriter(writer, os.Stdout)
	//}
	return pkg.NewZapLogger(
		encoder,
		writer,
		zap.NewAtomicLevelAt(zapcore.DebugLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCallerSkip(2),
		zap.Development(),
	)
}

// 日志自动切割，采用 lumberjack 实现的
func getLogWriter(fileName string) io.Writer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    10,    // 日志文件的最大大小（以MB为单位）
		MaxBackups: 5,     // 保留旧文件的最大个数
		MaxAge:     30,    // 保留旧文件的最大天数
		Compress:   false, // 是否压缩/归档旧文件
	}
	return lumberJackLogger
}
