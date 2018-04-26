package main

import "fmt"

func main(){
	switch "Medi"{
	case "Marcus":
		fmt.Println("Hello Marcus");
	case "Medhi":
		fmt.Println("Hello Medhi");
		fallthrough
	case "Hope":
		fmt.Println("Hello Hope");
		fallthrough
	case "Pearl":
		fmt.Println("Hello Pearl");
	}

}

/*
	When you do add a "fallthrough" the following
	case statement evaluates to true too.

	A default statement is not required. If none of the
	cases match the condition, it won't return anything.
*/