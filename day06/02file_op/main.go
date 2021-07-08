package main

import (
	"fmt"
	"io"
	"os"
)

//文件操作

//当err有值的时候吗，fileObj就是nil，nil不能调用Close()
func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
}

//使用这种写法
func f2() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
}

func f3() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("day06/02file_op/test.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//因为没有办法直接在文件中插入内容，所以要借助一个临时文件
	tmpFile, err := os.OpenFile("day06/02file_op/temp.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("creat temp file failed, err:%v\n", err)
		return
	}
	//读取文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err:%v\n", err)
		return
	}
	//写入临时文件
	_, _ = tmpFile.Write(ret[:n])
	//再写入要插入的内容
	var s []byte
	s = []byte{'c'}
	_, _ = tmpFile.Write(s)
	//紧接着把源文件的后续内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			_, _ = tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		_, _ = tmpFile.Write(x[:n])
	}
	//把源文件后续的也写入临时文件
	fileObj.Close()
	tmpFile.Close()
	_ = os.Rename("day06/02file_op/temp.txt", "day06/02file_op/test.txt")
}

func main() {
	//f1()
	f3()
}
