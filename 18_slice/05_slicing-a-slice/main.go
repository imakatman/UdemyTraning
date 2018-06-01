package main

import "fmt"

func main() {
	names := []string{
		"Hope Kim",
		"Pearl Kim",
		"Myong Contreres",
		"Sara Azoulay",
		"Leetal Eliyahu",
		"Twyla Garcia",
		"Carter Warren",
	}

	fmt.Println("[2:5]", names[2:5])
	fmt.Println("[2:5]", names[:5])
	fmt.Println("[2:5]", names[1:6])
}