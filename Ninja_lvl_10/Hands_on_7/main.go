package main

import "fmt"

func main() {
	c := make(chan int)
	genChan(c)
	//send
	for k := 0; k < 100; k++ {
		//v := range c {
		fmt.Println(k, <-c)
	}

}
func genChan(c chan int) {
	for i := 0; i < 10; i++ {
		go func() {
			foo(c)
		}()
	}
}

func foo(c chan<- int) {
	for i := 1; i < 11; i++ {
		c <- i
	}
}
