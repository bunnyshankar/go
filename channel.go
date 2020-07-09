package main

import "fmt"
import "sync"
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // creating channel of type int.
	wg.Add(2)
	go func() {
		i := <- ch // getting data from channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42 // write data to the channel
		wg.Done()
	}()
	wg.Wait()
}
