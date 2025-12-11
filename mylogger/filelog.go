package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level        LogLevel
	filePath     string // 日志文件保存的路径
	fileName     string // 日志文件保存的文件名
	fileObj      *os.File
	errorFileObj *os.File
	maxFileSize  int64
}

func NewFileLogger(levelStr, filePath, fileName string, maxFileSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	file := &FileLogger{
		Level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	err = file.initFile()
	if err != nil {
		panic(err)
	}
	return file
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}

	errorFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}

	f.fileObj = fileObj
	f.errorFileObj = errorFileObj
	return nil
}

// 日志开关
func (f *FileLogger) enalble(logLevel LogLevel) bool {
	return logLevel >= f.Level
}
func (f *FileLogger) chekaAndSplitFile(file *os.File) error {
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() >= f.maxFileSize {
		f.Close()

		fileName := path.Join(f.filePath, f.fileName)
		if file == f.errorFileObj {
			fileName += ".bck"
		}
		backupFileName := fmt.Sprintf("%s.bck%s", fileName, time.Now().Format("20060102150405"))
		os.Rename(fileName, backupFileName)
		// 创建新文件
		newFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		// 更新文件对象引用
		if file == f.fileObj {
			f.fileObj = newFile
		} else {
			f.errorFileObj = newFile
		}
	}
	return nil
}

func (f *FileLogger) checkFileSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		//	获取文件信息错误

		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFileLogger(file *os.File) (*os.File, error) {
	//需要切割的步骤
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info error \n")
		return nil, err
	}
	nowStr := time.Now().Format("20060102150405")
	//1，关闭当前的日志文件
	file.Close()
	//2，备份一下文件 Rename
	logName := path.Join(f.filePath, fileInfo.Name())
	backupFileName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	os.Rename(path.Join(f.filePath, f.fileName), backupFileName)
	//3，打开一个新的文件
	newFile, err := os.OpenFile(path.Join(f.filePath, f.fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new file error \n")
		return nil, err
	}
	//4，将打开的新日志文件对象 赋值给 fileObj
	return newFile, nil
}
func (f *FileLogger) log(lv LogLevel, formar string, args ...interface{}) {
	if f.enalble(lv) {
		msg := fmt.Sprintf(formar, args...)
		level := getLogString(lv)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		//fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, funcName, fileName, lineNo, msg)
		// f.chekaAndSplitFile(f.fileObj) //检查文件大小，是否创建新的文件
		if f.checkFileSize(f.fileObj) {
			newFile, err := f.splitFileLogger(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}

		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			if f.checkFileSize(f.errorFileObj) {
				fileObj, err := f.splitFileLogger(f.errorFileObj)
				if err != nil {
					return
				}
				f.errorFileObj = fileObj
			}
			fmt.Fprintf(f.errorFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), level, funcName, fileName, lineNo, msg)
		}
	}

}
func (f *FileLogger) Debug(msg string, args ...interface{}) {
	f.log(DEBUG, msg, args...)

}
func (f *FileLogger) Info(msg string, args ...interface{}) {
	f.log(INFO, msg, args...)
}
func (f *FileLogger) Warning(msg string, args ...interface{}) {
	f.log(WARNING, msg, args...)

}
func (f *FileLogger) Error(msg string, args ...interface{}) {
	f.log(ERROR, msg, args...)

}
func (f *FileLogger) Fatal(msg string, args ...interface{}) {
	f.log(FATAL, msg, args...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errorFileObj.Close()
}
