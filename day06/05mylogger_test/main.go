package main

import (
	"github.com/baqimeilaoban/studygo/day06/mylogger"
	"time"
)

//测试我们自己写的日志库

func main() {
	log := mylogger.NewLog("info")
	for true {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		time.Sleep(time.Second)
	}
}
