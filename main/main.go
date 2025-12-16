package main

import (
	"go-log/mylogger"
)

func main() {

	// log := mylogger.NewLog("error")
	log := mylogger.NewFileLogger("Info", "./", "mylog.log", 10*1024)

	// ticker := time.NewTicker(3 * time.Second)
	defer log.Close()
	// timeout := time.After(29 * time.Second)

	for {
		log.Debug("这是一条 Debug 日志")
		log.Debug("这是一条 Debug 日志")
		log.Info("这是一条 Info 日志")
		log.Warning("这是一条 Warning 日志")
		log.Fatal("这是一条 Fatal 日志")
		log.Error("这是一条 Error 日志,%d,%s,%f", 123, "abc", 3.14)
		log.Error("这是一条 Error 日志,%d,%s,%f", 123, "abc", 3.14)
		log.Error("这是一条 Error 日志,%d,%s,%f", 123, "abc", 3.14)
		log.Error("这是一条 Error 日志,%d,%s,%f", 123, "abc", 3.14)
		log.Error("这是一条 Error 日志,%d,%s,%f", 123, "abc", 3.14)
		log.Error("这是一条 Error 日志,%d,%s,%f", 123, "abc", 3.14)
	}

	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		log.Debug("这是一条 Debug 日志")
	// 		log.Info("这是一条 Info 日志")
	// 		log.Warning("这是一条 Warning 日志")
	// 		log.Fatal("这是一条 Fatal 日志")
	// 	case <-timeout:
	// 		log.Error("这是一条 Error 日志,%d,%s,%f", 123, "abc", 3.14)
	// 		return
	// 	}

	// }
}
