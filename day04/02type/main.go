package main

import "fmt"

//自定义类型和类别名

//type后面跟的是类型
type myInt int    //自定义类型
type youInt = int //类型别名

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	var m youInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
