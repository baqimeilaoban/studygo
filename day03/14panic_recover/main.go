package main

import "fmt"

//panic和recover

func funcA() {
	fmt.Println("a")
}

func funcB() {
	//刚刚打开数据库连接
	defer func() {
		err := recover() //基本上不会让使用
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重的错误")
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
