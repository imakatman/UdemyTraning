package main

import "fmt"

func main(){
	a := 24

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(b)
}