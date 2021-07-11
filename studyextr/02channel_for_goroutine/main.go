package main

import "fmt"

// printHello prints to stdout and puts an int on channel
func printHello(ch chan int) {
	fmt.Println("Hello from printHello")
	//send a value on channel
	ch <- 2
}

func main() {
	//make a channel, you need to use the make function to create channel
	//channels can also be buffered where you can specify size. eg: ch := make(chan int, 2)
	//that is out of the scope of this post.
	ch := make(chan int)
	//inline goroutine. Define a function and then call it.
	//write on a channel when done
	go func() {
		fmt.Println("Hello inline")
		//send a value on channel
		ch <- 1
	}()
	//call a function as goroutine
	go printHello(ch)
	fmt.Println("Hello from main")
	//get first value from channel
	//and assign to a variable to use this value later
	//here that to print it
	i := <-ch
	fmt.Println("Received", i)
	//get the second value from channel
	//do not assign it to a variable because we dont wat to use that
	<-ch
}
