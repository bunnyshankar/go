package main

import (
	"fmt"
)

func main() {
	s := "This is a string"
	s2 := "This is a 2nd string"
	fmt.Printf("%v, %T\n", s + s2, s + s2)
}
