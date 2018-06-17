package main

import "fmt"

type square struct{
	side float64
}

type shape interface {
	area() float64
}

func (s *square) area() float64{
	return s.side * s.side
}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	s := square{4}

	info(s) // this doesnt work because the area func is expecting a type pointer.
		// read quiver notes for more details.
}