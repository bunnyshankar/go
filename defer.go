package main

import "fmt"

func main() {
	//fmt.Println("First")
	//defer fmt.Println("Middle") //defer executes at the exit of the function.
	//fmt.Println("Last")

// defer follows lifo model. The last defer executes first
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Final")
}
