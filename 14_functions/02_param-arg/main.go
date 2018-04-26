package main

import "fmt"

func main(){
	greet("Hope")
	greet("Sara")
}

func greet(name string){
	fmt.Println("Hello", name);
}