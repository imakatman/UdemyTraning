package main

import "fmt"

func main(){
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos Dias",
		3: "Bongiorno",
	}

	fmt.Println(myGreeting)

	// val is the value of myGreetings[2]
	// exists is a bool that evaluates from checking whether the value of val exists
	if val, exists:=myGreeting[2]; exists{
		delete(myGreeting, 2)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
		fmt.Println(myGreeting)
	} else {
		fmt.Println("That value does not exist")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}
}