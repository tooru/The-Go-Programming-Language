package p05_uniq

import (
	"testing"
)

func assertEquals(t *testing.T, x, y []string) {
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

func TestRotate(t *testing.T) {
	assertEquals(t, nil, Uniq(nil))

	assertEquals(t, []string{"a"}, Uniq([]string{"a"}))
	assertEquals(t, []string{"a"}, Uniq([]string{"a", "a"}))

	assertEquals(t, []string{"a", "b"}, Uniq([]string{"a", "b"}))
	assertEquals(t, []string{"a", "b"}, Uniq([]string{"a", "b", "b"}))

	assertEquals(t, []string{"a", "b", "a"}, Uniq([]string{"a", "b", "a"}))
}
