package main

import "fmt"

func main(){
	greet("Hope", "Kim")
}

func greet(fname, lname string){
	fmt.Println("Hello", fname, lname);
}