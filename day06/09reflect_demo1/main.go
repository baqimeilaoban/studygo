package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v type kind:%v\n", v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 值的种类
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取整型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取整型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// reflectSetValue1 通过反射设置变量的值（值拷贝形式）
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		// 修改的是副本，reflect包会引发panic
		v.SetInt(200)
	}
}

// reflectSetValue2 通过反射设置变量的值（指针形式）
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// testIsNilOrIsValid 测试IsNil和IsValid
func testIsNilOrIsValid() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找出"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
	var c = Cat{}
	reflectType(c) // type:main.Cat

	// ValueOf
	reflectValue(a)
	// 设置值
	//reflectSetValue1(b)
	//fmt.Println(b)
	reflectSetValue2(&b)
	fmt.Println(b)

	testIsNilOrIsValid()
}
