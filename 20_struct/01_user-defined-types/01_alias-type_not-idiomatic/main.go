package main

import "fmt"

type foo int // not idiomatic

func main(){
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)
}