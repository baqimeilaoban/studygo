package main

import (
	"github.com/baqimeilaoban/studygo/day06/mylogger"
)

// log 声明一个全局的接口变量
var log mylogger.Logger

// 测试我们自己写的日志库
func main() {
	// 终端日志实例
	log = mylogger.NewConsoleLogger("Info")
	// 文件日志实例
	log = mylogger.NewFileLogger("info", "day06/05mylogger_test", "hjb.log", 10*1024*1024)
	for true {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条warning日志")
		id := 100010
		name := "理想"
		log.Error("这是一条Error日志.id:%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		//time.Sleep(2 * time.Second)
	}
}
