package main

import "fmt"

func main() {
	c := make(chan int)

	go func() { //by using the channel in a different go routine the main program can continue
		c <- 42 // this blocks this go routine
	}()
	fmt.Println(<-c)
}

//channels block until the send and receive can happen simultaniously.
//by using goroutines the send and receive can be done at the same time.
