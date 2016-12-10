package main

import (
    "fmt"
)

func main() {
    var _, _, c = toupleFunction()

    fmt.Println(c);
}

func toupleFunction() (string, string, string) {
    return "a", "b", "c"
}
