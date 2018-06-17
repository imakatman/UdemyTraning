package main

import "fmt"

func main(){
	name := "Hope"

	str, ok := name.(string); // syntax for assertion is <var>.(<type>)
	// this won't work because name isn't an interface

	if ok{
		fmt.Println(str)
	} else {
		fmt.Println("Value is not a string")
	}
}