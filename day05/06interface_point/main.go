package main

import "fmt"

//使用值接收者和指针接收者的区别

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

////使用值接收者实现了接口的所有方法
//func (c cat) move() {
//	fmt.Println("走猫步！")
//}
//
//func (c cat) eat(food string) {
//	fmt.Printf("猫吃%s\n", food)
//}

//使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步！")
}

func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	var a1 animal
	c1 := cat{ //cat
		name: "tom",
		feet: 4,
	}
	c2 := &cat{ //*cat
		name: "假老年",
		feet: 4,
	}
	//a1 = c1 //使用值接收时
	a1 = &c1 //使用指针接收时
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
