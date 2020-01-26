package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(4)
	if x != 28 {
		t.Error("Expected: 28, got:", x)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(4)
	if x != 28 {
		t.Error("Expected: 28, got:", x)
	}
}

func ExampleYears() {
	fmt.Println(Years(4))
	//Output:
	//28
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))
	//Output:
	//28
}
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(4)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(4)
	}
}
