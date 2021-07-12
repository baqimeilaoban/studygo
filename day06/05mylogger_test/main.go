package main

import (
	"github.com/baqimeilaoban/studygo/day06/mylogger"
	"time"
)

//测试我们自己写的日志库

func main() {
	log := mylogger.NewFileLogger("debug", "day06/05mylogger_test", "hjb.log", 10*1024*1024)
	for true {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		id := 1000
		name := "理想"
		log.Error("这是一条Error日志.id:%d,name:%s", id, name)
		time.Sleep(time.Second)
	}
}
