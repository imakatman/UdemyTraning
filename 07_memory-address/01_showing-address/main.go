package main

import "fmt"

func main(){
	a := "Hope Kim"
	fmt.Println("a is stored in/memory address is ", &a)
	b := &a
	fmt.Printf("%d\n", b)
}