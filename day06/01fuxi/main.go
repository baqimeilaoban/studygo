package main

import "fmt"

//类型断言

func main() {
	var a interface{} //定义一个空接口
	a = 100
	//如何判断a保存的值的具体类型是什么
	//类型断言
	//1.x.(T)
	if v, ok := a.(int8); ok {
		fmt.Println("猜对了，a是int8", v)
		return
	} else {
		fmt.Println("猜错了，不是int8")
	}

	//2.switch
	switch a.(type) {
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case int32:
		fmt.Println("int32")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("滚~")
	}
}
