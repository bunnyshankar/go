package main

import "fmt"


func main() {
	i  := 8 //2^3
	fmt.Println( i << 3) //left shift bit 2^3 * 2^3 = 2^6
	fmt.Println( i >> 3) // right shift bit = 2^3/2^3 = 2^0
}
