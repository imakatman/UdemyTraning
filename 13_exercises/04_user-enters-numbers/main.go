package main

import "fmt"

func main(){
	var big_no int32
	var small_no int32

	fmt.Print("Enter a big number: ")
	fmt.Scan(&big_no)

	fmt.Print("Enter a small number: ")
	fmt.Scan(&small_no)

	fmt.Println("The remainder of", big_no, "and", small_no, "is", big_no % small_no)
}

// Create a program that prints to the terminal asking for a user to enter a small number and a larger number. Print the remainder of the larger number divided by the smaller number.