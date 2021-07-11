package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// sleepRandomContext Function that does slow processing with a context. Note that context is the first argument
func sleepRandomContext(ctx context.Context, ch chan bool) {
	//Cleanup tasks
	//There are no contexts being created here
	//Hence, no canceling needed
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()
	//Make a channel
	sleepTimeChan := make(chan int)
	//Start a slow processing in a goroutine
	//Send a channel for communication
	go sleepRandom("sleepRandomContext", sleepTimeChan)
	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//If context expires, this case is selected
		//Free up resources that may no longer be needed because of aborting the work
		//Signal all the  goroutines that should stop work(use channels)
		//Usually you should send something on channel
		//wait for goroutines to exit and then return
		//Or, use wait groups instead of channels for synchronization
		fmt.Println("Time to return")
	case sleepTimeChan := <-sleepTimeChan:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("Slept for ", sleepTimeChan, "ms")
	}

}

// sleepRandom Slow function
func sleepRandom(fromFunction string, ch chan int) {
	//defer cleanup
	defer func() {
		fmt.Println(fromFunction, "sleepRandom complete")
	}()
	//Perform a slow task
	//For illustration purpose
	//Sleep here for random ms
	send := time.Now().UnixNano()
	r := rand.New(rand.NewSource(send))
	randomNumber := r.Intn(100)
	sleepTime := randomNumber + 100
	fmt.Println(fromFunction, "Start sleep for ", sleepTime, "ms")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleepTime, "ms")
	//write on the channel if it was passed in
	if ch != nil {
		ch <- sleepTime
	}
}

//A helper function, this can, in the real world do various things
//In this example, it is just calling one function
//Here, this could have just lived in main
func doWorkContext(ctx context.Context) {
	//Derive a timeout context from context with cancel
	//timeout in 150ms
	//All the contexts derived from this will returns in 150ms
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	//cancel to release resources once the function is complete
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()
	//Make channel and call context function
	//Can use wait groups as well for this particular case
	//As we do not use return value sent on channel
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)
	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//This case is selected when the passed in context notifies to stop work
		//In this example, it will be notified when main calls cancelFunction
		fmt.Println("doWorkContext: Time to Return")
	case <-ch:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("sleepRandomContext returned")
	}
}

func main() {
	//Make a background context
	ctx := context.Background()
	//Derive a context with cancel
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	//defer canceling so that all resources are free up
	//For this and the derived contexts
	defer func() {
		fmt.Println("Main defer: canceling context")
		cancelFunction()
	}()
	//Cancel context after a random time
	//This cancels the request after a random timeout
	//If this happens, all the contexts derived from this should return
	go func() {
		sleepRandom("main", nil)
		cancelFunction()
		fmt.Println("Main Sleep complete, canceling context")
	}()
	//Do work
	doWorkContext(ctxWithCancel)
}
