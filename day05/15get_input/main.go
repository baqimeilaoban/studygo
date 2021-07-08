package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入时如果有空格

func useScan() {
	var s string
	fmt.Println("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的内容是:%s", s)
}

func useBufio() {
	var s string
	fmt.Println("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是:%s\n", s)
}

func main() {
	//useScan()
	//useBufio()

	//写日志
	fmt.Fprintln(os.Stdout, "这是一条日志记录！")
	fileObj, _ := os.OpenFile("day05/15get_input/test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(fileObj, "这是一条日志记录")
}
