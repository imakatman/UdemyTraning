package main

import "fmt"

var g string = "cowboy"

func main(){
	a := 0
	b := "golang"
	c := 4.17
	d := true

	// %v is a verb that prints the value in a default format
	fmt.Printf("%v \n", a);
	fmt.Printf("%v \n", b);
	fmt.Printf("%v \n", c);
	fmt.Printf("%v \n", d);
}