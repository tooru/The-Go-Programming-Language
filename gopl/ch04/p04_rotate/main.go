package p04_rotate

// 2-pass?

func Rotate(s []int, r int) {
	t := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		t[(i+r)%len(s)] = s[i]
	}
	for i := 0; i < len(s); i++ {
		s[i] = t[i]
	}
}
