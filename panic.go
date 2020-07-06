package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("this is deferred") // at end of function, defer executes before the panic.
	panic("something happend") //exit status 2 
	fmt.Println("end")
}

