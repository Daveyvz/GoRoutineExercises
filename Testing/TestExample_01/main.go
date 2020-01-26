package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", myAdd(2, 3))
	fmt.Println("4 + 7 = ", myAdd(4, 7))
	fmt.Println("8 + 6 = ", myAdd(8, 6))
}

//myAdd adds the given int values together and returns that value.
func myAdd(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
