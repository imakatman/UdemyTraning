package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func main(){
	// notes are just to help with using proper language
	p1 := person{"Hope", "Kim", 25} // creating a value of type person and assigning it to variable p1
	p2 := person{"Carter", "Warren", 28}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)

	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.age)
}