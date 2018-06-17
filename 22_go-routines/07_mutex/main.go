package main

import (
	"time"
	"sync"
	"fmt"
	"math/rand"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func incrementor(s string){
	for i := 0; i < 20; i++{
		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond) // this is moved to the top because you don't want this to be locked, otherwise
		mutex.Lock() // This locks any other threads from trying to access any of the code that follows
		// The following commented out code was used for the purposes of demonstrating a race condition
		//x := counter
		//x++
		//counter = x

		// No one has access to the counter variable while it's locked so we can write to the global variable
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()  // This unlocks any other threads from trying to access any of the code that follows

		//some people wrap braces around the code to make it easier to identify what code is being locked
		//{
		//	mutex.Lock()
		//	counter++
		//	fmt.Println(s, i, "Counter:", counter)
		//	mutex.Unlock()
		//}
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
}