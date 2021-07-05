package main

import "fmt"

//结构体嵌套
type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address //匿名嵌套
	workPlace
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "周林",
		age:  9000,
		address: address{
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	//fmt.Println(p1.city) //可以直接p1.city获取城市信息，现在自己结构体照这个字段，找不到就去匿名嵌套的结构体中查找该字段
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
