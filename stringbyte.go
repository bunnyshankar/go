package main

import (
	"fmt"
)

func main() {
	s := "This is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
