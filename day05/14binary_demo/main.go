package main

import "fmt"

//二进制实际用途

const (
	eat   int = 4 //0100
	sleep int = 2 //0010
	play  int = 1 //0001
)

//111
//左边的1表示吃饭
//中间的1表示睡觉
//右边的1表示打豆豆
func color(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	color(eat | play)         //0101
	color(eat | sleep | play) //0111
}
