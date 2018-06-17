package main

import (
	"sync"
	"fmt"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	// This is a race group because the first two go routines are trying to add to a global variable.

	go func() {
		wg.Add(1)
		for i := 0; i < 20; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 20; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func(){
		wg.Wait()
		close(c)
	}()

	for n := range c{
		fmt.Println(n)
	}
}