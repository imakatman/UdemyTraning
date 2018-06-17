package main

import "fmt"

func main(){
	var name interface{} = "Hope"
	// name is of a concrete type, empty interface, and its assigned a value type string "Hope"

	str, ok := name.(string) // this will only work if the value assigned to name is a string

	if ok{
		fmt.Printf("%T \n", str)
	} else {
		fmt.Println("Value is not a string")
	}
}