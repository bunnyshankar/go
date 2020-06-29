package main

import "fmt"

func main() {
//	arrays := [3]int{67,78,91} //no of elements and type of array and values of the array
	arrays := [...]int{67,78,91} //another way of declaring
	fmt.Printf("%v\n", arrays)
	var students [3]string
	students[0]="shankar"
	fmt.Printf("%v\n", students)
	fmt.Printf("%v\n", len(students)) //size of an array

// array of arrays
	var indexMatrix [3][3]int
	indexMatrix[0] = [3]int{1,0,0}
	indexMatrix[1] = [3]int{0,1,0}
	indexMatrix[2] = [3]int{0,0,1}
	fmt.Printf("%v\n", indexMatrix)
	fmt.Printf("%v\n", len(indexMatrix))
//2nd way for array of arrays
	var indexM [3][3]int = [3][3]int{[...]int{1,0,0}, [...]int{0,1,0}, [...]int{0,0,1}}
	fmt.Printf("%v\n", indexM)

//pointers in array
	a := [3]int{1, 2, 3}
	b := a // copy of a is now in b
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := &a //pointer to 'a' array
	c[1] = 50
	fmt.Println(a)
	fmt.Println(c)
}
