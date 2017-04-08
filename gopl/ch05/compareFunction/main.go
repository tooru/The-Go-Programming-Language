package main

func one() int { return 1; }
func two() int { return 2; }

func main() {
    // compile error
    //one == two // invalid operation: one == two (func can only be compared to nil)
}
