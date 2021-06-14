package main

import "fmt"

//map

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil) //true 还没在内存中开辟空间
	m1 = make(map[string]int, 10) //要估算好改map容量，避免在程序运行期间动态扩容
	m1["理想"] = 9000
	m1["姬无命"] = 35
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	//约定成俗用OK接收返回的布尔值
	fmt.Println(m1["娜扎"]) //如果不存在这个key，则返回这个对应值类型的零值
	value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此人")
	}else {
		fmt.Println(value)
	}

	//map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//只遍历k
	for k := range m1 {
		fmt.Println(k)
	}

	//只遍历v
	for _, v := range m1 {
		fmt.Println(v)
	}

	//删除
	delete(m1, "姬无命")
	fmt.Println(m1)
	delete(m1, "沙河") //删除不存在的key时，不
}
