package main

import (
	"math"
	"fmt"
)

type circle struct{
	radius float64
}

type shape interface{
	area() float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func info(s shape){
	fmt.Println(s.area())
}

func main(){
	c := circle{40}

	/*
	All following functions work
	*/
	fmt.Println(shape.area(c))
	info(c)
}