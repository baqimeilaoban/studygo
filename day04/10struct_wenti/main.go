package main

import (
	"fmt"
)

//结构体遇到的问题

//1.myInt(100)是个啥？
type myInt int

type person struct {
	name string
	age  int
}

func (m myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	//声明一个int32类型的变量x，它的值是10
	//方法1
	var x1 int32
	x1 = 10
	fmt.Printf("type:%T value:%d\n", x1, x1)
	//方法2
	var x2 int32 = 10
	fmt.Printf("type:%T value:%d\n", x2, x2)
	//方法3
	var x3 = int32(10)
	fmt.Printf("type:%T value:%d\n", x3, x3)
	//方法4
	x4 := int32(10)
	fmt.Printf("type:%T value:%d\n", x4, x4)

	//声明一个myInt类型的变量m，它的值是100
	//方法1
	var m1 myInt
	m1 = 100
	fmt.Printf("type:%T value:%d\n", m1, m1)
	//方法2
	var m2 myInt = 100
	fmt.Printf("type:%T value:%d\n", m2, m2)
	//方法3
	var m3 = myInt(100) //强制类型转换
	fmt.Printf("type:%T value:%d\n", m3, m3)
	//方法4
	m4 := myInt(100)
	fmt.Printf("type:%T value:%d\n", m4, m4)

	m := myInt(100)
	m.hello()

	//问题2：结构体初始化
	//方法1
	var p person //声明一个person类型的变量p
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)
	//方法2：键值对初始化
	var p2 = person{
		name: "官话",
		age:  18,
	}
	fmt.Println(p2)
	//方法3：值列表初始化
	var p3 = person{
		"理想",
		100,
	}
	fmt.Println(p3)
}

//问题3：为啥要有构造函数
func newPerson(name string, age int) person {
	//别人调用我，我能给她一个person类型的变量
	return person{
		name: name,
		age:  age,
	}
}

// func newPerson(name string, age int) *person {
// 	return &person{
// 		name: name,
// 		age: age,
// 	}
// }
