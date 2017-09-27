package main

func main() {
	type A struct {
		X int
	}
	type B struct {
		X int
		Y int
	}
	type C struct {
		A
		B
		Y int
	}

	var c C

	//c.X = 1 // CompileError ambiguous selector c.X
	c.A.X = 1
	c.B.X = 1
	c.Y = 1
}
