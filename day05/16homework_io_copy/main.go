package main

import (
	"fmt"
	"io"
	"os"
)

//借助io.Copy实现一个拷贝文件函数

func copyFile(dstName string, srcName string) (written int64, err error) {
	//以只读的方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", srcName, err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("close %s failed, err:%v\n", file, err)
			return
		}
	}(src)
	//以写|创建的方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dstName, err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("close %s failed, err:%v\n", file, err)
		}
	}(dst)
	return io.Copy(dst, src)
}

func main() {
	_, err := copyFile("day05/16homework_io_copy/dst.txt", "day05/16homework_io_copy/src.txt")
	if err != nil {
		fmt.Printf("copy file failed, err:%v\n", err)
		return
	}
	fmt.Println("copy done")
}
