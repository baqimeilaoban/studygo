package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	// 从字符串中解析出对应的数据
	str := "1000"
	//ret1 := int64(str)
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parseint failed, err:%v\n", err)
		return
	}
	fmt.Printf("the type:%T the value:%d\n", ret1, ret1)
	// 字符串->int类型
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("the type:%T the value:%d\n", retInt, retInt)
	// 把数字转换成字符串类型
	i := 97
	//ret2 := string(i) // "a"
	ret2 := fmt.Sprintf("%d", i) // "97"
	fmt.Printf("%#v\n", ret2)
	// 数字->字符串
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v\n", ret3)
	// 字符串->布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)
	// 字符串->浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)
}
