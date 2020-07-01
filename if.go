package main

import "fmt"

func main() {
	statePopulations := map[string]int{
                "califormia":   12121,
                "Texas":        12121,
                "Mayfair":      43534,
                "oregon":       44524,
        }
	if pop, ok := statePopulations["oregon"]; ok {
		fmt.Println(pop)
	}
	number := 50
	guess := -5
	if guess < number {
		fmt.Println("Too Low")
	}
	if guess > number {
		fmt.Println("Too High")
	}
	if guess == number {
		fmt.Println("Yo Got it!")
	}
	fmt.Println(number<=guess, number >= guess, number!=guess)
	//logical operators
	if guess < 1 || guess > 100 {
		fmt.Println("Guess should be between 1 to 100")
	}
	if guess >=1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too Low")
		}
		if guess > number {
			fmt.Println("Too High")
		}
		if guess == number {
			fmt.Println("Yo Got it!")
		}
		fmt.Println(number<=guess, number >= guess, number!=guess)
	}
	// if else code
        if guess < 1 || guess > 100 {
                fmt.Println("Guess should be between 1 to 100")
        } else {
                if guess < number {
                        fmt.Println("Too Low")
                }
                if guess > number {
                        fmt.Println("Too High")
                }
                if guess == number {
                        fmt.Println("Yo Got it!")
                }
                fmt.Println(number<=guess, number >= guess, number!=guess)
        }
	// multiple if conditions
	fmt.Println("Multiple If conditions")
        if guess < 1 {
                fmt.Println("Guess should be greater than 1")
	} else if  guess > 100 {
		fmt.Println("Guess should be less than 100")
        } else {
                if guess < number {
                        fmt.Println("Too Low")
                }
                if guess > number {
                        fmt.Println("Too High")
                }
                if guess == number {
                        fmt.Println("Yo Got it!")
                }
                fmt.Println(number<=guess, number >= guess, number!=guess)
        }


}
