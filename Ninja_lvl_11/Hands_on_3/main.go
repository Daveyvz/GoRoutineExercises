package main

import "fmt"

type customErr struct{}

func (c customErr) Error() string {
	return "there was an error"
}

func main() {
	c := customErr{}

	foo(c)
}

func foo(e error) {
	fmt.Println(e.Error())
}
