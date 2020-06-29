package main

import (
	"fmt"
)
// Emurated constants
const (
	_ = iota //ignore first value by assigning to a blank identifier
	a
	b
	c
)
const (
	d = iota //resets the value of iota to 0
)
func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
