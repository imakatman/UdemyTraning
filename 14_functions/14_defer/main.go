package main

import "fmt"

func world(){
	fmt.Println("world")
}

func the(){
	fmt.Println("the")
}

func hello(){
	fmt.Println("hello ")
}

func main(){
	hello()
	defer world()
	the()
}

// defer is a keyword that will defer a function
// until the function it's in will exit