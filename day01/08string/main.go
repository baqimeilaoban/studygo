package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// \本来具有特殊含义的，我应该告诉程序我写的\就是一个单纯的\
	path := "\"\\Users\\huangjianbo\\go\\src\\github.com\\baqimeilaoban\\studygo\\day01\\08string\""
	fmt.Println(path)

	s := "I'm ok"
	fmt.Println(s)

	// 多行的字符串
	s2 := `
		世情薄，
		人情恶，
		雨送黄昏花易落。
	`
	fmt.Println(s2)

	s3 := `\Users\huangjianbo\go\src\github.com\baqimeilaoban\studygo\day01\08string`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "理想"
	word := "大帅比"
	ss := name + word
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, word) //将其拼接后返回成一个字符串变量
	fmt.Println(ss1)
	fmt.Printf("%s%s\n", name, word) //直接返回拼接后的字符串

	// 分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret) //输出数组形式

	// 包含
	fmt.Println(strings.Contains(ss, "理性")) //输出false

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcdb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
