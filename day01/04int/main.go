package main

import "fmt"

// 整形

func main()  {
	var i1 = 100
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制转换为二进制 
	fmt.Printf("%o\n", i1) //把十进制转换为八进制
	fmt.Printf("%x\n", i1) //把十进制转换为十六进制
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量的类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型的变量
	i4 := int8(9) //明确指定int8类型，否则就是默认的int类型
	fmt.Printf("%T\n", i4)
}