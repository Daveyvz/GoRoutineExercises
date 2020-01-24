package main

import "fmt"

func main() {
	c := make(chan int, 1) //the buffer tells the channel that 1 value may be reserved into it.

	c <- 42 //there is only 1 value assigned to the channel therefore it doesn't block
	//c <- 43  	< this can't be added because then there would be more than 1 value in the channel and it would block

	fmt.Println(<-c) //the first value is retrieved
	//fmt.Println(<-c) //the second value is retrieved
}
