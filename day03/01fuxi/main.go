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
	var ages3 = [...]int{1:100, 99:200}
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
	y := x //把x的值拷贝了一份给y
	y[1] = 200 //修改的是副本y，并不影响x
 	fmt.Println(x) //[1 2 3]
	f1(x)
	fmt.Println(x) //[1 2 3]
}

func f1(a [3]int)  {
	//Go语言中的函数传递的都是值（Ctrl+C Ctrl+V）
	a[1] = 100 //此处修改的是副本的值
}
