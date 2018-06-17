package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		//close() will close a channel, you are supposed to close a channel once it is exhausted
		close(c)
	}()

	/*
	this for loop waits for a value to be added to the channel and prints it.
	range is used to range over the values that are passed to c.
	once it's exhausted it will exit.
	 */
	for n := range c {
		fmt.Println(n)
	}

	/*
	There is no need to add a time wait because the code is blocked by the above
	for loop until c is exhausted
	 */
}