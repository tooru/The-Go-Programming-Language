package main

import (
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

func main() {
	s := "abcdefgあ𩸽🇯🇵"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%d\t%v\n", i, r, size, utf16.IsSurrogate(r))
		i += size
	}
}
