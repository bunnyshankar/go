package main

import (
	"fmt"
	"strconv"
)

func main() {
	var k int = 42
	var j float32
	j = float32(k) //convert into float32
	var m int = int(j) //convert into inti
	var s string = strconv.Itoa(m) //convert int to string Itoa is integer to ascii
	fmt.Printf("%v, %T", k, k)
	fmt.Printf("%v, %T", j, j)
	fmt.Printf("%v, %T", m, m)
	fmt.Printf("%v, %T", s, s)
}
