package main

import "fmt"

func main(){
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}

// A "break" statement will exit you out of the for loop