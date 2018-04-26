package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "This is a string inside of the inner scope"
		fmt.Println(y)
	}
	// y is not accessible in the outer scope because y was declared inside of the inner scope
	//fmt.Println(y)
}
