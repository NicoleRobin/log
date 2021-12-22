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
	defaultLogger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}

func SetLevel(level Level) {
	if defaultLogger == nil {
		panic("defaultLogger is nil")
	}
	if zapLevel, ok := levelToZapLevel[level]; ok {
		defaultLogger.Core().Enabled(zapLevel)
	} else {
		panic("unknow log level")
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
