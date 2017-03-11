package main

import (
    "fmt"
)

func main() {
    fmt.Printf("%#v\n", nil)
    // fmt.Printf("%v\n", nil[:]) // SyntaxError use of untyped nil
    // fmt.Printf("%v\n", nil[:0]) // SyntaxError use of untyped nil

    var s []int

    fmt.Printf("%#v\n", s)
    fmt.Printf("%#v\n", s[:])
    fmt.Printf("%#v\n", s[:0])

    s = make([]int, 0)
    fmt.Printf("%#v\n", s)
    fmt.Printf("%#v\n", s[:])
    fmt.Printf("%#v\n", s[:0])
}
