package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	for i := 0; i < 80; i++{
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), " Cap: ", cap(mySlice), " Item: ", i);
	}

}