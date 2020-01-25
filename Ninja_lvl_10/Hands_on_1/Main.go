package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	e := make(chan int, 1)

	e <- 43
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Println(<-e)
}
