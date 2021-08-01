package main

import (
	"encoding/json"
	"fmt"
)

// 反射

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"zhoulin","age":9000}`
	var stu1 student
	err := json.Unmarshal([]byte(str), &stu1)
	if err != nil {
		fmt.Printf("Parsing json is failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", stu1)
}
