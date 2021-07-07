package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件

//读取文件方法1：循环读取文件(自己设立一次读多少)
func readFromFile1() {
	fileObj, err := os.Open("day05/12file_open/main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//记得关闭文件
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("close file failed, err:%v\n", err)
			return
		}
	}(fileObj)
	//读文件
	//var tmp = make([]byte, 128) //指定读的长度
	var tmp [128]byte
	for true {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

//利用bufio这个包读取文件(一行一行读取)
func readFromFileByBufio() {
	fileObj, err := os.Open("day05/12file_open/main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//关闭文件
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Printf("close file failed, err:%v\n", err)
			return
		}
	}(fileObj)
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for true {
		line, err := reader.ReadString('\n')
		//读到文件末尾时退出
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func readFileByIoutil() {
	ret, err := ioutil.ReadFile("day05/12file_open/main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	readFromFile1()
	readFromFileByBufio()
	readFileByIoutil()
}
