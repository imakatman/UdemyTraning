package main

import "fmt"

func main(){
	for i:=0; i<=100; i++{
		if i % 2 == 0{
			fmt.Println(i)
		}
	}
}

//Print all of the even numbers between 0 and 100.