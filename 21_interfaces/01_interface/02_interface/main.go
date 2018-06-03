package main

import "fmt"

type Square struct{
	side float64
}

// Attaching method area() to Square struct
func (z Square) area() float64{
	return z.side * z.side
}

type Shape interface{
	/*
	Anything that implements this method
	signature, area() float64, is implementing
	the Shape interface. *NOTE: area() doent take a
	parameter so in order to use the Shape interface
	you cant pass a parameter to area().

	So type Square has this method signature attached to
	 it, so it is implementing the Shape interface.
	*/
	area() float64
}

func info(z Shape){
	fmt.Println(z)
	fmt.Println(z.area())
}

func main(){
	s := Square{10}
	info(s)
}



