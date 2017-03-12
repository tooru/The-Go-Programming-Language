package main

import (
    "fmt"
)

const (
    KB = 1000
    MB = KB * 1000
    GB = MB * 1000
    TB = GB * 1000
    PB = TB * 1000
    EB = PB * 1000
    ZB = EB * 1000
    YB = ZB * 1000
)    


func main() {
    fmt.Printf("%d", KB)
    fmt.Printf("%d", MB)
    fmt.Printf("%d", GB)
    fmt.Printf("%d", TB)
    fmt.Printf("%d", PB)
    fmt.Printf("%d", EB)
    //fmt.Printf("%d", ZB) // overflow int unprintable?
    //fmt.Printf("%d", YB) // overflow int unprintable?
}
