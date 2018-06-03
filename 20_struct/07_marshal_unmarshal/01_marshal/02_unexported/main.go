package main

import (
	"fmt"
	"encoding/json"
)

type Person struct{
	first string
	last string
	age int
}

func main(){
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs)) // this will print out an empty string because the fields in Person arent exported
}