package main

import "fmt"

//变量的作用域

var x = 100 //定义一个全局变量

func f1() {
	//x := 10
	name := "理想"
	//函数中查找变量的顺序
	//1.现在函数内部查找
	//2.找不到就往函数外面查找，一直找到全局
	fmt.Println(x, name)
}

func main() {
	f1()
	//fmt.Println(name) //函数内部定义的变量只能在该函数内部使用
	//语句块作用域
	if i := 10; i < 18 {
		fmt.Println("乖乖上学")
	}
	//fmt.Println(i)
}
