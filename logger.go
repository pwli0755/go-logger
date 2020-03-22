package logger

import (
	"fmt"
	"sync"
)


var (
	logger ILogger
	loggerMu sync.RWMutex
	loggers   = make(map[string]ILogger)
)

func InitLogger(typ string, config ...interface{}) error {
	loggerMu.RLock()
	loggerImp, ok := loggers[typ]
	loggerMu.RUnlock()
	if !ok{
		return fmt.Errorf("logger: unknown logger type %s(forgotten register?)", typ)
	}
	logger = loggerImp
	return nil
}

func Register(typ string, loggerImp ILogger) {
	loggerMu.Lock()
	defer loggerMu.Unlock()
	if loggerImp == nil {
		panic("logger: Register logger is nil")
	}
	if _, dup := loggers[typ]; dup {
		panic("logger: Register called twice for logger type " + typ)
	}
	loggers[typ] = loggerImp
}

func Debug(format string, args ...interface{})  {
	logger.Debug(format, args...)
}
func Trace(format string, args ...interface{})  {
	logger.Trace(format, args...)
}
func Info(format string, args ...interface{})  {
	logger.Info(format, args...)
}
func Warn(format string, args ...interface{})  {
	logger.Warn(format, args...)
}
func Error(format string, args ...interface{})  {
	logger.Error(format, args...)
}
func Fatal(format string, args ...interface{})  {
	logger.Fatal(format, args...)
}
