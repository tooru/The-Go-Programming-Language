package p04_rotate

import (
	"testing"
)

func assertEquals(t *testing.T, x, y []int) {
	if len(x) != len(y) {
		t.Errorf("%v %v", x, y)
		return
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			t.Errorf("%v %v", x, y)
			return
		}
	}
}

func rotate(s []int, r int) []int {
	Rotate(s, r)

	return s
}

func TestRotate(t *testing.T) {
	assertEquals(t, nil, rotate(nil, 0))

	assertEquals(t, []int{0}, rotate([]int{0}, 0))
	assertEquals(t, []int{0}, rotate([]int{0}, 1))

	assertEquals(t, []int{0, 1}, rotate([]int{0, 1}, 0))
	assertEquals(t, []int{1, 0}, rotate([]int{0, 1}, 1))
	assertEquals(t, []int{0, 1}, rotate([]int{0, 1}, 2))

	assertEquals(t, []int{0, 1, 2}, rotate([]int{0, 1, 2}, 0))
	assertEquals(t, []int{2, 0, 1}, rotate([]int{0, 1, 2}, 1))
	assertEquals(t, []int{1, 2, 0}, rotate([]int{0, 1, 2}, 2))
	assertEquals(t, []int{0, 1, 2}, rotate([]int{0, 1, 2}, 3))
}
