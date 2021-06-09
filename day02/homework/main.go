package main

import "fmt"

// 单行注释

/*
多行注释
*/

// Go语言函数外的语句必须以关键字开头
var name = "娜扎"
var age int

//age = 19
const (
	num = 100
)

// 如果要编译可执行文件，必要要有main包和main函数入口(入口函数)
// main函数是入口函数，没有参数也没有返回值
func main() {
	// 函数内部定义的变量必须使用
	var isOK = true
	fmt.Println(isOK)

	// if
	if age > 18 {
		fmt.Println("成年了")
	}else if age > 7{
		fmt.Println("上小学")
	}else {
		fmt.Println("最快乐的时光")
	}

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for.2
	var i = 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	// for.3
	var j = 0
	for ; j < 10; {
		fmt.Println(j)
		j++
	}
	
	//for.4
	fmt.Println("无限循环")

	//for.5
	s := "hello"
	fmt.Println(len(s)) //5
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// 哑元变量用法，不想用到的直接扔给他
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}

	// 八进制
	var n1 = 0777
	// 十六进制
	var n2 = 0xff
	fmt.Println(n1, n2)
	fmt.Printf("%o\n", n1)
	fmt.Printf("%x\n", n2)
}
