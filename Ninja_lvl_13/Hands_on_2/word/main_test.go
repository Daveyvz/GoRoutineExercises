package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	x := Count("Davey is a legend")
	if x != 4 {
		t.Error("Expected: 4, got:", x)
	}
}

func ExampleCount() {
	fmt.Println(Count("Davey is a legend"))
	//Output:
	//4
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Davey is a legend")
	}
}
