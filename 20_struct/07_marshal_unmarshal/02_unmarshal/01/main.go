package main

import (
	"fmt"
	"encoding/json"
)

type Person struct{
	First string
	Last string
	Age int
}

func main(){
	var p1 Person
	fmt.Println(p1.First) // emptystring
	fmt.Println(p1.Last) // emptystring
	fmt.Println(p1.Age) // 0

	bs := []byte(`{"First": "James","Last": "Bond", "Age": 20 }`)

	json.Unmarshal(bs, &p1) // Unmarshal this byte slice and assign its value to p1

	fmt.Println("--------------------------------------------")
	fmt.Println(p1.First) // James
	fmt.Println(p1.Last) // Bond
	fmt.Println(p1.Age) // 20
}