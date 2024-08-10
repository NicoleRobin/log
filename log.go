package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level int

const (
	LevelNil Level = iota
	LevelDebug
	LevelWarn
	LevelInfo
	LevelError
)

var levelToZapLevel = map[Level]zapcore.Level{
	LevelDebug: zapcore.DebugLevel,
	LevelInfo:  zapcore.InfoLevel,
	LevelWarn:  zapcore.WarnLevel,
	LevelError: zapcore.ErrorLevel,
}

var defaultLogger *zap.Logger

func init() {
	var err error
	developmentConfig := zap.NewDevelopmentConfig()
	defaultLogger, err = developmentConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func SetLevel(level Level) {
	if defaultLogger == nil {
		panic("defaultLogger is nil")
	}
	if zapLevel, ok := levelToZapLevel[level]; ok {
		var err error
		developmentConfig := zap.NewDevelopmentConfig()
		developmentConfig.Level = zap.NewAtomicLevelAt(zapLevel)
		defaultLogger, err = developmentConfig.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	} else {
		panic("unknown log level")
	}
}

func Debug(msg string, args ...interface{}) {
	defaultLogger.Debug(fmt.Sprintf(msg, args...))
}

func Info(msg string, args ...interface{}) {
	defaultLogger.Info(fmt.Sprintf(msg, args...))
}

func Warn(msg string, args ...interface{}) {
	defaultLogger.Warn(fmt.Sprintf(msg, args...))
}
func Error(msg string, args ...interface{}) {
	defaultLogger.Error(fmt.Sprintf(msg, args...))
}
