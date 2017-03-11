package main

func main() {
    size2 := [2]int{1,2}
    size3 := [3]int{1,2,3}

    // cannot convert size3 (type [3]int) to type [2]int
    // cannot convert size2 (type [2]int) to type [3]int
    [2]int(size3)
    [3]int(size2)
}
