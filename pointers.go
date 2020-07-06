package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a // b holds address memory of a. '*' represents pointer to a datatype.
	fmt.Println(a,b)
	fmt.Println(&a, b) // to check memory address of both a and b
	fmt.Println(a, *b)  // to check value stored in the memory address of a pointer.
	*b = 50 //assigning values to pointers by deferencing it.
	fmt.Println(a, *b)
	c := [3]int{1,3,4}
	d := &c[1]
	e := &c[2]
	fmt.Printf("%v %p %p\n", c,d,e) 
	// lttle more on pointers
	var ms *myStruct
	fmt.Println(ms) // This returns a nil value. Important to check with pointers 
	ms = new(myStruct) // initializing the object
	ms.foo = 123
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
