package main

import "fmt"


func main() {
	i  := 10 //binary representation 1010
	j  := 3 // binart representation 0011
	fmt.Println( i & j) //0010 and
	fmt.Println( i | j) //1011 or
	fmt.Println( i ^ j) //1001 not
	fmt.Println( i &^ j) // 0100 xor
}
