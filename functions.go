package main

import "fmt"

func main() {
	sayHello("This is calling function")
	greeting("Hello", "Shjankar")
	s := sum(1, 2, 3, 4, 5, 6)
	fmt.Println("The sum is",s)
	d, err := divide(5.0, 1.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	//Anonymous function
	func() {
		fmt.Println("Anonymous")
	}()
}


func sayHello(msg string) {
	fmt.Println(msg)
}
func greeting(msg, name string) { //every parameter as a string
	fmt.Println(msg,name,"!!")
}
func sum(values ...int) int { //slices of the values passed to this function returns int value variadic parameters type. It has to be last in the paramterlist
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result //returns a value to calling function
}
func divide (a float64, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a/b, nil //no error so retur nil for error type
}
