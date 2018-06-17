package main

import (
	"runtime"
	"fmt"
	"time"
	"sync"
)

// init() code runs before anything else that is being executed in this script
func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	/*
	As of Go 1.5, this code is unnecessary but what is doing is
	telling Go to use all of the Cores available in the OS to
	run the following code. You need more than one Core to execute
	parallelism
	*/
}

var wg sync.WaitGroup

func foo(){
	for i := 0; i < 45; i++{
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar(){
	for i := 0; i < 45; i++{
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}