package main

import "fmt"

func main(){
	var age interface{} = 25

	fmt.Println(age.(int) + 5)
}