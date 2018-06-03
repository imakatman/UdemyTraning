package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

type alien struct{
	first string
	last string
	age int
}

/*
(p person) is the receiver of this func.
any value of the type person will have access
to the method fullname().
it's not idiomatic to say "this", or "self",
check the notes to read more.
 */
func (p person) fullName() string{
	return p.first + p.last
}

func main(){
	p1 := person{"Hope", "Kim", 25}
	p2 := person{"Ryan", "Pardeiro", 37}
	a1 := alien{"Yolandi", "Visser", 37}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	//fmt.Println(a1.fullName()) // not going to work because not of type person
}