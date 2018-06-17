package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int

func incrementor(s string){
	for i := 0; i < 20; i++{
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond) // this ensures that code runs concurrently
		counter = x
		fmt.Println(s, i, "Counter: ", counter)
	}

	wg.Done()
}

func main(){
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	/*
	this will always be a different result (i.e. 20, 21, 22)
	because theres some randomness about when they are writing
	*/
	fmt.Println("Final Counter: ", counter)
}