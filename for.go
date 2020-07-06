package main

import "fmt"

func main() {
	for i :=0; i < 5; i++ {
		fmt.Println(i)
	}
	for i :=0; i < 5; i = i + 2 {
		fmt.Println(i)
	}
	for i, j := 0, 0; i < 5 || j <3 ; i, j = i+1,  j+1 {
		fmt.Println(i,j)
	}
	i := 1
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	j := 1
	for j < 5 {
		fmt.Println(j)
		j++
	}
	//infinite loops make sure you have break for a condition
	for {
		fmt.Println(j)
		j++
		if j == 20 {
			break //break the loop
		}
	}
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue //Dont execute further statements and continue the loop
		}
		fmt.Println(i)
	}
	firstloop: //label for the outermost forloop
	for i := 0; i<5; i++ {
		for j := 0;  j < 3; j++ {
			fmt.Println(i*j)
			if i * j >= 6 {
				break firstloop  //break the outmost forloop with the label
			}
		}
	}
	s := []int{1, 2, 3}
	for k, v := range s { //for looping through a slice we use range 
		fmt.Println(k,v)
	}
}

