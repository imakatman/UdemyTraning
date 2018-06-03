package main

import "fmt"

type creature struct{
	first string
	last string
	age int
}

type person struct{
	creature
	first string
	last string
	age int
	race string
	occupation string
}

func main(){
	c1 := creature{"Chase", "Kim", 12}
	p1 := person{
		creature: creature{
			first: "Hope",
			last: "Kim",
			age: 25,
		},
		first: "Somang",
		last: "Kim",
		age: 30,
		race: "Asian",
		occupation: "Web Developer",
	}

	fmt.Println(c1.first, c1.last, c1.age)
	fmt.Println(p1.first, p1.last, p1.age) // will print the first, last, and age properties outside of the creature type
}