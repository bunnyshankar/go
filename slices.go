package main

import "fmt"

func main() {
	slices := []int{67,78,91} //another way of declaring
	fmt.Printf("%v\n", slices)
	fmt.Printf("%v\n", len(slices)) //size of an slice
	fmt.Printf("%v\n", cap(slices)) //Capacity of an underlying array of a slice

// array of slices
//	var indexMatrix [][]int
//	indexMatrix[0] = []int{1,0,0}
//	indexMatrix[1] = []int{0,1,0}
//	indexMatrix[2] = []int{0,0,1}
//	fmt.Printf("%v\n", indexMatrix)
//	fmt.Printf("%v\n", len(indexMatrix))
//2nd way for array of slices
//	var indexM [][]int = [][]int{[]int{1,0,0}, []int{0,1,0}, []int{0,0,1}}
//	fmt.Printf("%v\n", indexM)

//References in slices
	m := []int{1, 2, 3}
	n := m // In slices, its reference. Reference of a is now copied to b.
	n[1] = 5
	fmt.Println(m)
	fmt.Println(n)

//Some other workouts in slices
	a := []int{1, 2,3 ,4, 5, 6, 7, 8, 9, 10}
	b := a[:] //slice all elements
	c := a[3:] //slice from 4th element to end
	d := a[:6] //slice first 6 elements
	e := a[3:6]//slice the 4th, 5th and 6th elements
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

// make function 
//	mk := make([]int, 3, 100) //type of slice, initial size of slice  and capacity of slice
	mk := make([]int, 3) //to initialize a slice: type of slice, initial size of slice  and capacity of slice
	mk = append(mk, 1) //to add an element to a slice
	mk = append(mk, []int{3,4,41}...) // concatinating a slice into another with spread operator(...)
	fmt.Println(mk)
	fmt.Println(len(mk))
	fmt.Println(cap(mk))
	remove := append(a[:2], a[4:]...) //to remove elements from a slice
	fmt.Println(remove)
}
