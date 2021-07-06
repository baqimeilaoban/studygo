package main

import (
	"fmt"

	calc "github.com/baqimeilaoban/studygo/day05/10calc"
)

var x = 100

const pi = 3.14

func main() {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}

func init() {
	fmt.Println("自动执行！")
	fmt.Println(x, pi)
}
