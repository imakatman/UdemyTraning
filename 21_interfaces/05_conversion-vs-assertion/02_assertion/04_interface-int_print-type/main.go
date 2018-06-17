package main

import "fmt"

func main(){
	var age interface{} = 25

	fmt.Printf("%T \n", age)
}