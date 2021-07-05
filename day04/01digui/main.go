package main

import "fmt"

//递归：自己调用自己
//递归适合处理那种问题相同\问题的规模越来越小的场景
//递归一定要有一个明确的退出条件

//上台阶：n个台阶，一次可以走一步，也可以走两步，有多少步可以走

func taijie(n uint64) uint64 {
	if n == 1 {
		//如果只有一个台阶，就只有一种走法
		return 1
	}
	if n == 2 {
		//如果两个台阶，就有两种走法
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

//5! = 5*4*3*2*1 = 5*4!

//计算n的阶乘

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	n := 5
	ret := f(uint64(n))
	fmt.Printf("%d的阶乘为：%d\n", n, ret)
	sum := taijie(uint64(n))
	fmt.Printf("%d个台阶有%d种走法！", n, sum)
}
