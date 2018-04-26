package main

import "fmt"

func findGreatest(ln ...float64) float64{
	var g float64
	for _, n := range ln{
		if n > g{
			g = n
		}
	}
	return g
}

func main(){
	fmt.Println(findGreatest(40, 1, 7, 89, 50))
}