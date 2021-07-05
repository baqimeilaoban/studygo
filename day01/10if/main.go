package main

import "fmt"

// if条件判断

func main() {
	age := 19
	if age > 18 { //如果age大于18，就执行这个{}中的代码
		fmt.Println("澳门首家线上赌场开业了")
	} else { //否则就执行这个{}中的代码
		fmt.Println("改写暑假作业了")
	}

	// 多个判断
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

	// 作用域 age1变量此时只在if条件判断语句中生效，减少程序内存的占用
	if age1 := 19; age1 > 18 {
		fmt.Println("澳门首家线上赌场开业啦")
	} else {
		fmt.Println("该写暑假作业了")
	}
	//fmt.Println(age1) //在这里找不到age1
}
