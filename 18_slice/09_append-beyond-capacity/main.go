package main

import "fmt"

func main(){
	names := make([]string, 4, 10)

	names[0] = "Hope Kim"
	names[1] = "Pearl Iwao"
	names[2] = "Myong Contreres"
	names[3] = "Sara Azoulay"
	names = append(names, "Leetal Eliyahu")
	names = append(names, "Carter Warren")
	names = append(names, "Twyla Garcia")

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))
}