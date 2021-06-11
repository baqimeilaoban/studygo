package main

import "fmt"

// goto

func main() {
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break //跳出内层for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break //跳出外层for循环
		}
	}

	// goto + label实现跳出多层循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	XX:
		fmt.Println("over")
}
