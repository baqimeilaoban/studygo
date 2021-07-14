package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Address 试验用到的数据
type Address struct {
	City    string
	Country string
}

// 序列化后数据存放的路径
var filePath string

// encode 将数据序列号后写入文件中
func encode() {
	pa := &Address{
		City:    "Chengdu",
		Country: "China",
	}
	// 打开文件，不存在的时候新建
	flie, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer flie.Close()
	// encode后写入这个文件中
	enc := gob.NewEncoder(flie)
	enc.Encode(pa)
}

// decode 从文件中读取数据并反序列化
func decode() *Address {
	file, _ := os.Open(filePath)
	defer file.Close()
	var pa Address
	// decode操作
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	return &pa
}

func main() {
	filePath = "studyextr/encode/serialization_dserialization/address.gob"
	encode()
	pa := decode()
	fmt.Println(*pa)
}
