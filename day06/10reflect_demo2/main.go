package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	// student struct
	fmt.Println(t.Name(), t.Kind())
	// 通过for循环遍历结构体的所以字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if scoreFiled, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreFiled.Name, scoreFiled.Index, scoreFiled.Type, scoreFiled.Tag.Get("json"))
	}
}
