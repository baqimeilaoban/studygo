package main

import "fmt"

//闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2进行包装
func lixiang(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

func main() {
	ret := lixiang(f2, 100, 200) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	ret()
}
