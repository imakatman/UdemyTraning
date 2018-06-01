package main

import "fmt"

func main() {
	// this will create a nil reference and you would never do this with a map
	// has no length because its nil
	//var ages map[string]int

	// These are both valid ways to initialize a map
	var ages = make(map[string]int)
	var ages2 = map[string]int{} // shorthand/ composite literal

	ages["Hope Kim"] = 24
	ages["Pearl Iwao"] = 35


	var ages3 = map[string]int{
		"Hope Kim":  24,
		"Pearl Iwao":  34,
	}

	fmt.Println(ages)
	fmt.Println(ages2)
	fmt.Println(ages3)
	fmt.Println(len(ages3))

	ages3["Hope Kim"] = 25

	fmt.Println(ages3)

	var greetings = map[int]string{
		0: "Hello",
		1: "Bonjour",
		2: "Hola",
	}

	fmt.Println(greetings)

	delete(greetings, 1)

	fmt.Println(greetings)
}