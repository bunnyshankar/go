package main

import "fmt"

func main() {
	switch 2 {
	case 1:
		fmt.Println("onr")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
//range of options
	switch 5 {
	case 1, 5, 10:
		fmt.Println("One, five, ten")
	case 2, 6, 9:
		fmt.Println("two, six, nine")
	default:
		fmt.Println("No match")
	}
//Tag syntax
        switch i:=2+3; i {
        case 1, 5, 10:
                fmt.Println("TAG One, five, ten")
        case 2, 6, 9:
                fmt.Println("two, six, nine")
        default:
                fmt.Println("No match")
        }
//TagLess Syntax
	j := 20
        switch {
        case j <= 10:
                fmt.Println("less than or equal ten")
		fallthrough //keyword used for executing the next case even though it doesnt satisfy.
        case j == 20:
                fmt.Println("Its  20")
		break // to break a case not executing further
		fmt.Println("This will not be printed")
        default:
                fmt.Println("greater than 20")
        }

}
