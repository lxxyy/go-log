package mylogger

import (
	"fmt"
	"time"
)

// 日志开关
func (l *Logger) enalble(logLevel LogLevel) bool {
	return logLevel >= l.Level
}

func (l *Logger) log(lv LogLevel, formar string, args ...interface{}) {
	msg := fmt.Sprintf(formar, args...)
	level := getLogString(lv)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, funcName, fileName, lineNo, msg)

}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.enalble(DEBUG) {
		l.log(DEBUG, msg, args...)
	}

}
func (l *Logger) Info(msg string, args ...interface{}) {
	if l.enalble(INFO) {
		l.log(INFO, msg, args...)
	}

}
func (l *Logger) Warning(msg string, args ...interface{}) {
	if l.enalble(WARNING) {
		l.log(WARNING, msg, args...)
	}

}
func (l *Logger) Error(msg string, args ...interface{}) {
	if l.enalble(ERROR) {
		l.log(INFO, msg, args...)
	}

}
func (l *Logger) Fatal(msg string, args ...interface{}) {
	if l.enalble(FATAL) {
		l.log(INFO, msg, args...)
	}
}
