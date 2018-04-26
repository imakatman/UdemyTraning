package main

import "fmt"

func main(){
	switch "Jenny" {
	case "Tim", "Hope":
		fmt.Println("Hello Tim or Hope")
	case "Marcus", "Medhi":
		fmt.Println("Hello Marcus or Medhi")
	case "Jenny", "Pearl":
		fmt.Println("Hello Jenny or Pearl")
	}
}

/*
	You can add multiple evaluations in a case statement.
	But you can't add duplicate evaluations in different
	case statements.
*/