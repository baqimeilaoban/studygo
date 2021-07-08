package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

//写文件方法1
func writeDemo1() {
	fileObj, err := os.OpenFile("day05/13file_write/xxx.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	fileObj.Write([]byte("zhoulin mengbi le!\n"))
	fileObj.WriteString("周林解释不了！\n")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("fileObj close failed, err:%v\n", err)
		}
	}(fileObj)
}

//写文件方法2
func writeDemo2() {
	fileObj, err := os.OpenFile("day05/13file_write/xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file faile, err:%v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("close file failed, err:%v\n", err)
			return
		}
	}(fileObj)
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	_, err = wr.WriteString("hello沙河\n") //写在缓冲区中
	if err != nil {
		fmt.Printf("write to cache failed, err:%v\n", err)
		return
	}
	err = wr.Flush()
	if err != nil {
		fmt.Printf("write to file failed, err:%v\n", err)
		return
	}
}

//文件写法3
func writeDemo3() {
	str := "黄建波，早上好\n"
	err := ioutil.WriteFile("day05/13file_write/xxx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write to file failed, err:%v\n", err)
		return
	}
}

func main() {
	writeDemo1()
	writeDemo2()
	writeDemo3()
}
