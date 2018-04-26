package main

import "fmt"

var x int = 25

func main() {
	fmt.Println(x)
	foo()
	y := 7
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
