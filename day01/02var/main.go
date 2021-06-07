package main

import "fmt"

// Go语言中推荐使用驼峰式的命名方法
// var student_name string //下划线
var studentName string //小驼峰，推荐
// var StudentName string //大驼峰

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	//var heihei string 局部声明后必须使用
	name = "理想"
	age = 18
	isOk = true

	// Go语言中变量声明了必须使用，不使用就编译不过去
	fmt.Print(isOk)             //在终端中输出要打印的内容
	fmt.Printf("name:%s", name) // %s:占位符 使用name这个变量值去替换占位符
	fmt.Println(age)            //打印完指定内容之后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "王海波"
	fmt.Println(s1)
	// 类型推导，根据值判断该变量是什么类型
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明，只能在函数里面使用
	s3 := "哈哈哈"
	fmt.Println(s3)
	// s1 := "10" // 同一个作用域({})不能重复声明同名的变量
}
