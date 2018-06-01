package main

import "fmt"

func main(){
	aSlice := make([]int, 1)
	fmt.Println(aSlice[0])
	aSlice[0] = 7
	fmt.Println(aSlice[0])
	aSlice[0]++ // adds one to the value of the integer at index 0 in aSlice
	fmt.Println(aSlice[0])
}