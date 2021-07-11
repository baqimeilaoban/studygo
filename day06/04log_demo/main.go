package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//log demo

func main() {
	fileObj, err := os.OpenFile("day06/04log_demo/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	log.SetOutput(fileObj) //os.Stdout:往终端中输出
	for true {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 2)
	}
}
