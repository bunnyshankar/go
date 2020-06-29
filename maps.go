package main

import "fmt"

func main() {
	//Key of Strings of type int
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"califormia":	12121,
		"Texas":	12121,
		"Mayfair":	43534,
		"oregon":	44524,
	}
	fmt.Println(statePopulations)
	statePopulations["Ohio"] = 321312
	fmt.Println(statePopulations)
	delete(statePopulations, "Ohio") // delete from a map 
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"]) // printing a value of a map
	_, ok := statePopulations["Texas"] // ,ok syntax to check if key is existing in a map
	fmt.Println(ok)
}

