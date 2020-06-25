package main

import "fmt"

// var block

var (
	actor string = "Shankar"
	age int = 40
	movie string = "telus"
)

func main() {
       //type 1 
//       var i int
//        i = 42
      
       //type 2 controlled way of variable type
  //     var j float32 = 27
      
       // type 3 - go decides the type
       k := 99.
       
        
	fmt.Printf("%v, %T", k, k)
}
