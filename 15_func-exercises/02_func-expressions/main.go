package main

import "fmt"

func main(){
	half := func(x int) (float64, bool){
		f := float64(x)
		return f / 2.0, x % 2 == 0
	}

	h1, even1 := half(5)
	h2, even2 := half(8)
	fmt.Println(h1, even1)
	fmt.Println(h2, even2)
}