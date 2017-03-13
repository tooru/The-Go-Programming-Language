package main

import (
    "fmt"
)

func reverse(a *[5]int) {
    for i, j := 0, 4; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
    }
}

func main() {
    a := [...]int{1,2,3,4,5}
    reverse(&a)
    fmt.Printf("%v\n", a)
}
