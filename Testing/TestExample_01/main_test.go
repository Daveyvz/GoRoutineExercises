package main

import "testing"

func TestMyAdd(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	//tabletest
	tests := []test{
		test{[]int{25, 25}, 50},
		test{[]int{14, 16}, 30},
		test{[]int{14, 16, 20}, 50},
		test{[]int{4, 6}, 10},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := myAdd(v.data...)
		if x != v.answer {
			t.Error("Expected:", v.answer, "got:", x)
		}
	}
}
