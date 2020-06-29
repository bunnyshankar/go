package main

import (
	"fmt"
)

const a string = "this is for package level"
func main() {
	const a int = 14  //typed constants int. constant a is getting shadowed.
	const b float32 = 1.2 //typed constants float32
	const c string = "This is constant" //typed constants string
	const d bool = true //typed constants boolean
	const e = 45 // during compile time, compiler decides the datatype associated with it.
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", b + e, b + e)
	
}
