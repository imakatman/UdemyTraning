package main

import "fmt"

type foo int // not idiomatic

func main(){
	var myAge foo
	myAge = 25
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 44
	fmt.Printf("%T %v \n", yourAge, yourAge)

	//fmt.Println(myAge + yourAge)
	// wont work because myAge and yourAge are of different types

	fmt.Println(int(myAge)+ yourAge)
	// will work because youve converted myAge to a type of int
}