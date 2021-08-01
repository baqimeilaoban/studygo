package main

import (
	"encoding/json"
	"fmt"
)

// person 结构体
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"周林","age":9000}`
	var p person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Printf("Parsing json is failed, err:%v\n", err)
		return
	}
	fmt.Println(p.Name, p.Age)
}
