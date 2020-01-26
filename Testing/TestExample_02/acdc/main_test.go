package acdc

import (
	"fmt"
	"testing"
)

//ExampleSum functions as both an example on the documentation and as a test.
func ExampleSum() {
	fmt.Println(Sum(2, 3))
	//Output:
	//5
}

func TestSum(t *testing.T) {
	x := Sum(2, 3)
	if x != 5 {
		t.Error("Expected: 5, Got:", x)
	}
}
