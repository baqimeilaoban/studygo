package main

import (
	"encoding/json"
	"fmt"
)

//复习结构体

type tmp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

type point struct {
	X int `json:"周林"`
	Y int `json:"保德路"`
}

func sum(x int, y int) (ret int) {
	ret = x + y
	return ret
}

//构造函数（结构体变量的）函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//方法
//接收者使用对应类型首字母的小写
//指定了接收者之后，只有接收者这个类型的变量才能调用这个方法
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是%s\n", p.name, str)
}

func (p person) guonian() {
	p.age++ //此处p是p1的副本，改的是副本
}

//指针接收者
//1.需要修改结构体变量的值时需要使用指针接收者
//2.结构体本身比较大，拷贝的内存开销比较大时也要使用指针接收者
//3.保持一致性：如果一个方法使用了指针接收者，其他方法为了统一也要使用指针接收者
func (p *person) zhenguonian() {
	p.age++ //
}

//结构体嵌套
type addr struct {
	province string
	city     string
}

type student struct {
	name string
	addr //匿名嵌套别的结构体，就使用类型名做名称
}

func main() {
	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	var b = tmp{
		10,
		20,
	}
	fmt.Println(b)

	var x int
	y := int8(10)
	fmt.Println(x, y)

	var p1 person //结构体实例化
	p1.name = "周林"
	p1.age = 16
	p2 := person{"周林", 100} //结构体实例化
	fmt.Println(p1, p2)
	//调用构造函数生成person类型变量
	p4 := newPerson("娜扎", 19)
	fmt.Println(p4)

	p1.dream("做个咸鱼")
	p2.dream("学好Go语言")
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.zhenguonian()
	fmt.Println(p1.age)

	//序列化
	po1 := point{100, 200}
	b1, err := json.Marshal(po1)
	//如果出错了
	if err != nil {
		fmt.Printf("Marshal failed, err:%v\n", err)
	}
	fmt.Println(string(b1))

	//反序列化：由字符串 --> Go中的结构体变量
	str := `{"周林":100, "保德路":200}`
	//造一个结构体变量，准备接收反序列化的值
	var po2 point
	err = json.Unmarshal([]byte(str), &po2)
	if err != nil {
		fmt.Printf("Unmarshal failed, err:%v\n", err)
	}
	fmt.Println(po2)

	//map
	m1 := map[int64]string{
		10081: "哈哈哈",
		10010: "嘿嘿嘿",
		10000: "呵呵呵",
	}
	name := m1[10010]
	fmt.Println(name)
 
	name2, ok := m1[100000] //ok=true表示有这个key，ok=false表示没有这个key
	fmt.Println(name2, ok)
}
