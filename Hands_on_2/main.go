package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("say something")
}
func saySomething(h human) {
	h.speak()
}
func main() {
	p := person{
		first: "davey",
		last:  "van zetten",
		age:   23,
	}

	//saySomething(p)				// doens't work because you can't use p as a parameter (it doesn't implicitly implement interface human)
	saySomething(&p) // does work because it gives a pointer to the adress where p is stored and that is what speak needs so p has speak and therefore is a human
}
