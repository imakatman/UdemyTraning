package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(2 * time.Millisecond))
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()
}

func main(){
	/*
	A go routine is capable of running concurrently with other funcs.
	 If you don't add a weight group, and you try to run two any go routines,
	 main() will exit before the go routines end.
	*/
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
