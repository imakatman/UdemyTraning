package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan int) // channel that takes an int

	/*
	The analogy for channels is relay-racing
	The first go routine is putting variable i into c and it doesnt run again until a value is taken off the channel
	The second go routine is printing the value that is stored into c
	Once it prints, the value is taken off the channel and the first go routine may commence
	*/
	go func(){
		for i := 0; i < 20; i++{
			c <- i // notation is here is take i and put it onto the channel, c
		}
	}()

	go func(){
		for{
			fmt.Println(<-c) // from the channel, give me whatever's there
		}
	}()

	/*
	Instead of using a WaitGroup which requires more code, we just have the program
	wait a second before exiting which is more than enough time for the go routines
	to execute.
	*/
	time.Sleep(time.Second)
}