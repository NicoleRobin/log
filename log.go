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

func newConfig() zap.Config {
	encoderConf := zap.NewProductionEncoderConfig()
	// encoderConf.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	encoderConf.EncodeTime = zapcore.RFC3339TimeEncoder

	logConf := zap.NewProductionConfig()
	logConf.EncoderConfig = encoderConf
	logConf.Encoding = "console"
	logConf.OutputPaths = []string{"stdout"}
	logConf.ErrorOutputPaths = []string{"stdout"}
	return logConf
}

func newLogger(config zap.Config) *zap.Logger {
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return logger
}

func init() {
	logConfig := newConfig()
	logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	defaultLogger = newLogger(logConfig)
}

func SetLevel(level Level) {
	if defaultLogger == nil {
		panic("defaultLogger is nil")
	}
	if zapLevel, ok := levelToZapLevel[level]; ok {
		logConfig := newConfig()
		logConfig.Level = zap.NewAtomicLevelAt(zapLevel)
		defaultLogger = newLogger(logConfig)
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
