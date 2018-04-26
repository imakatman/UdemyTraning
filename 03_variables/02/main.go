package main

import "fmt"

var g string = "cowboy"

func main(){
	var a int
	var b string
	var c float64
	var d bool

	// %v is a verb that prints the value in a default format
	fmt.Printf("%v \n", a);
	fmt.Printf("%v \n", b);
	fmt.Printf("%v \n", c);
	fmt.Printf("%v \n", d);
}