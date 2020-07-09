package main

import "fmt"
import "sync"
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // creating channel of type int.
	wg.Add(2)
	go func(ch <-chan int) { //receiver only channel
		i := <- ch // getting data from channel
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // writer only channel
		ch <- 42 // write data to the channel
		wg.Done()
	}(ch)
	wg.Wait()
}
