package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go foo(c)

	//receive
	bar(c) // this will wait until the foo goroutine catches up and the send and receive interject (because it blocks this goroutine)

	fmt.Println("about to exit")
}

//send
func foo(c chan<- int) {
	c <- 42
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

/*c := make(chan<- int, 2)   can only send ints so the println cant read from it because the channel isnt allowed to receive
//c := make(<-chan int, 2)	can only receive ints so you cant send ints to this channel.
//c := make(chan int, 2)
//c <- 42
//c <- 43

fmt.Println(<-c)
fmt.Println(<-c)
fmt.Println("-------")
fmt.Printf("%T\n", c) */
