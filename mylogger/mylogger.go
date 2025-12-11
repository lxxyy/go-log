package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

// 日志级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// 日志对象
type Logger struct {
	Level LogLevel
}

// 构造日志对象
func NewLog(levelStr string) *Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &Logger{Level: level}
}

// 字符串转日志级别
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
func getInfo(n int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[0]
	return
}
