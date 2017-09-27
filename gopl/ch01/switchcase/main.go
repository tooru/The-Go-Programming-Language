package main

import (
	"fmt"
	"os"
)

func main() {
	len := len(os.Args)

	fmt.Printf("%d\n", len)

	switch {
	case len < 3:
		fmt.Println("len < 3")
	case len < 2:
		fmt.Println("len < 2")
	default:
		fmt.Println("default")
		//fallthrough
	}
	fmt.Println("switch")
}
