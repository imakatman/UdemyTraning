package main

import "fmt"

func main(){
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos Dias",
		3: "Bongiorno",
	}

	for key, value := range myGreeting{
		fmt.Println(key, "- ", value)
	}
}