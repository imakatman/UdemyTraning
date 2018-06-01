package main

import "fmt"

func main(){
	mySlice := []string{"Hello", "Hola",}
	mySlice2 := []string{"Bye", "Adios",}

	mySlice = append(mySlice, mySlice2...) // to append a slice to a slice you must use the ... syntax

	mySlice3 := []int{1,3,5,}

	mySlice = append(mySlice, mySlice3...)

	fmt.Println(mySlice)
}