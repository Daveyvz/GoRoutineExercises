package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	xi := []int{2, 3, 4, 5, 6}
	x := CenteredAvg(xi)
	if x != 4 {
		t.Error("Expected: 2, got:", x)
	}
}

func ExampleCenteredAvg() {
	xi := []int{2, 3, 4, 5, 6}
	fmt.Println(CenteredAvg(xi))
	//Output:
	//4
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}
