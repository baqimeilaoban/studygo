package main

import "fmt"

//copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 //赋值
	//var a3 []int //nil，没空间
	var a3 = make([]int, 3) //切片
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//将a1中索引为1的3这个元素删掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)
	fmt.Println(cap(a1))

	x1 := [...]int{1, 3, 5} //数组
	s1 := x1[:]             //切片
	fmt.Println(s1, len(s1), cap(s1))
	//1.切片不保存具体的值
	//2.切片对应一个底层数组
	//3.底层数组都是占用一块连续的内存
	s1 = append(s1[:1], s1[2:]...) //修改了底层数组
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1) //[1 5 5]
}
