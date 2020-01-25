package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

var error = fmt.Errorf("custom err")

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs := toJSON(p1)
	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) []byte {
	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		fmt.Println(error)
	}
	return bs
}
