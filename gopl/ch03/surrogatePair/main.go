package main

import (
    "fmt"
    "unicode/utf8"
    "unicode/utf16"
)

func main() {
    s := "abcdefgあ𩸽🇯🇵"

    for i := 0; i < len(s); {
        r, size := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%d\t%c\t%d\t%v\n", i, r, size, utf16.IsSurrogate(r))
        i += size
    }
}
