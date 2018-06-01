package main

import (
	"fmt"
)

func main() {
	student := []string{} // default value for length and cap is 0
	students := [][]string{}
	fmt.Println(len(student))
	fmt.Println(cap(student))
	fmt.Println(students)
	fmt.Println(student == nil)

	student = append(student, "Carter")
	fmt.Println(student)
}
