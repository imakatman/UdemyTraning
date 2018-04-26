package main

import "fmt"

func main(){
	fmt.Println(greet("Hope", "Kim"))
}

func greet(fname, lname string) string{
	// Sprint means "String Print", so its printing it to a string.
	return fmt.Sprint(fname, lname);
}

// When returning a value you must enter the type after the parameter operators ()