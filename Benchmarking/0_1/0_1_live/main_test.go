package live

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Davey")
	if s != "Davey" {
		t.Error("expected: Davey, got:", s)
	}
}

func ExampleGreet() {
	fmt.Println("Davey")
	//Output:
	//Davey
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Davey")
	}
}
