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

func (c creature) fullName() string{
	return c.first + c.last
}

func (p person) fullName() string{
	return p.first + p.last
}

func main(){
	c1 := creature{"Chase", "Kim", 12}
	p1 := person{
		creature: creature{
			first: "Hope",
			last: "Kim",
			age: 5,
		},
		first: "Somang",
		last: "Kim",
		age: 25,
		race: "Asian",
		occupation: "Web Developer",
	}

	fmt.Println(c1.fullName()) // ChaseKim
	fmt.Println(p1.fullName()) // SomangKim
	fmt.Println(p1.creature.fullName()) // HopeKim
}