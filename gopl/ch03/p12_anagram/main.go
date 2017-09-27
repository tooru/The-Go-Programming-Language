package main

import (
	"fmt"
	"os"
)

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		if a[i] != b[j] {
			return false
		}
	}
	return true
}

func main() {
	a := os.Args[1]
	b := os.Args[2]

	fmt.Printf("%s %s => %v\n", a, b, isAnagram(a, b))
}
