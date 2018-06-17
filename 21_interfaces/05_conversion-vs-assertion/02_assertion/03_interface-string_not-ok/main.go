package main

import "fmt"

func main(){
	var name interface{} = 17

	str, ok := name.(string) // this won't work because the value of name is of type int

	if ok{
		fmt.Printf("%T \n", str)
	} else {
		fmt.Println("value is not of type string")
	}
}