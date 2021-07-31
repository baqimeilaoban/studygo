package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("day06/07file_demo/main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//1.文件对象的类型
	fmt.Printf("%T\n", fileObj)
	//2.获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return
	}
	fmt.Printf("文件名是%v，文件大小是:%d\n", fileInfo.Name(), fileInfo.Size())
}
