package logger

const (
	LevelDebug = iota
	LevelTrace
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	SplitByTime = iota
	SplitBySize
)

func getLevelText(level int) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelTrace:
		return "TRACE"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	}
	return "UNKNOWN"
}

func getLogLevel(level string) int {
	switch level {
	case "debug":
		return LevelDebug
	case "trace":
		return LevelTrace
	case "info":
		return LevelInfo
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	}
	return LevelDebug
}