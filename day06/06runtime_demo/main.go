package main

import (
	"fmt"
	"path"
	"runtime"
)

//runtime.Caller

func f() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) //06runtime_demo/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line) //11
}

func f1() {
	f()
}

func main() {
	f1()
}
