package main

import "fmt"

//复习

func main() {
	var name = "理想"
	fmt.Println(name)
	var ages [30]int
	ages = [30]int{1, 3, 5}
	fmt.Println(ages)
	var ages2 = [...]int{1, 3, 4, 5}
	fmt.Println(ages2)
	var ages3 = [...]int{1: 100, 99: 200}
	fmt.Println(ages3)
	//二维数组
	var a1 = [...][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a1)
	//多维数组只有最外层可以添加...

	//数组是值类型
	x := [3]int{1, 2, 3}
	y := x         //把x的值拷贝了一份给y
	y[1] = 200     //修改的是副本y，并不影响x
	fmt.Println(x) //[1 2 3]
	f1(x)
	fmt.Println(x) //[1 2 3]

	//切片(slice)：
	var s1 []int //没有分配内存，==nil
	fmt.Println(s1)
	fmt.Println(s1 == nil)
	s1 = []int{1, 2, 3}
	fmt.Println(s1)
	//make初始化 分配内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2)
	s3 := make([]int, 0, 4)
	fmt.Println(s3 == nil)

	x1 := []int{1, 2, 3}
	x2 := x1
	x3 := make([]int, 3)
	copy(x3, x1)
	fmt.Println(x2)
	x2[1] = 100
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)

	var c1 []int
	c1 = append(c1, 1) //自动初始化切片
	fmt.Println(c1)

	//指针
	//Go语言中的指针只能读不能改，不能修改指针变量指向的地址
	addr := "沙河"
	addp := &addr
	fmt.Println(addp) //addr指向的内存地址
	fmt.Printf("%T\n", addp)
	addrV := *addp //根据内存地址取值
	fmt.Println(addrV)

	//map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["jiaxiang"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"]) //如果key不存在返回的是value对应的零值
	//如果返回值是布尔类型，我们通常使用OK去接收
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有姬无命这个人")
	} else {
		fmt.Println("姬无命的分数是：", score)
	}
	delete(m1, "lixiang") //删除的key不存在就什么也不干
	delete(m1, "jiaxiang")
	fmt.Println(m1)
	fmt.Println(m1 == nil) //已经开辟了内存
}

func f1(a [3]int) {
	//Go语言中的函数传递的都是值（Ctrl+C Ctrl+V）
	a[1] = 100 //此处修改的是副本的值
}
