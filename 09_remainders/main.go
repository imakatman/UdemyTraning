package main

import "fmt"

func main(){
	for i := 0; i < 100; i++ {
		if i % 2 == 1 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
	}
}