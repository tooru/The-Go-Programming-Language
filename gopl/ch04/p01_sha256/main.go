package main

import (
    "fmt"
    "crypto/sha256"
)

var pc int

func popCount(x byte) int {
    return int((x & 0x80) >> 7 + 
               (x & 0x40) >> 6 + 
               (x & 0x20) >> 5 + 
               (x & 0x10) >> 4 + 
               (x & 0x08) >> 3 + 
               (x & 0x04) >> 2 + 
               (x & 0x02) >> 1 + 
               (x & 0x01))
}

func diff(c1, c2 *[32]byte) int {
    r := 0
    for i := 0; i < len(c1); i++ {
        r += popCount(c1[i] ^ c2[i])
    }
    return r
}

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))


    fmt.Printf("%d\n", diff(&c1, &c2))
}
