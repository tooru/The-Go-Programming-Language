package main

import (
    "fmt"
    "gopl/ch02/p01_kelvin"
)

func main() {
    fmt.Printf("%s\n", tempconv.AbsoluteZeroC)
    fmt.Printf("%s\n", tempconv.FreezingC)
    fmt.Printf("%s\n", tempconv.BoilingC)

    fmt.Printf("%s\n", tempconv.AbsoluteZeroK)
    fmt.Printf("%s\n", tempconv.FreezingK)
    fmt.Printf("%s\n", tempconv.BoilingK)
}
