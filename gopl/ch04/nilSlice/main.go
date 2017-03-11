package main

import (
    "fmt"
)

func main() {
    fmt.Printf("%#v\n", nil)
    // fmt.Printf("%v\n", nil[:]) // SyntaxError use of untyped nil
    // fmt.Printf("%v\n", nil[:0]) // SyntaxError use of untyped nil

    var s []int

    fmt.Printf("%#v %v\n", s, s == nil)
    s = s[:]
    fmt.Printf("%#v %v\n", s, s == nil)
    s = s[:0]
    fmt.Printf("%#v %v\n", s, s == nil)

    s = make([]int, 0)

    fmt.Printf("%#v %v\n", s, s == nil)
    s = s[:]
    fmt.Printf("%#v %v\n", s, s == nil)
    s = s[:0]
    fmt.Printf("%#v %v\n", s, s == nil)

    // go v1.7.3
    // <nil>
    // []int(nil) true
    // []int(nil) true
    // []int(nil) true
    // []int{} false
    // []int{} false
    // []int{} false
}
