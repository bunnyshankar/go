package main

import "fmt"
import "time"
import "sync"

//var wg = sync.WaitGroup{} // syncs all the routines
func main() {
	go sayHello() //with go we are starting a go routine which spans a blue thread and execute.
	time.Sleep(100 * time.Millisecond) //to delay to see the execution of thread
}

func sayHello() {
	fmt.Println("Helo")
}
