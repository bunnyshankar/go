package main

import (
	"fmt"
)

func main() {
	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", string(s[2]), s[2]) //to print 3rd character from string it has to be converted to string from uint
}
