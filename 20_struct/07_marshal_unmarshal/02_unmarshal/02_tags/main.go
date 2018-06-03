package main

import (
	"fmt"
	"encoding/json"
)

type Person struct{
	First string
	Last string
	Age int `json:"wisdom score"` // if something comes in as wisdom score, store it in age
}

func main(){
	var p1 Person
	var p2 Person
	fmt.Println(p1.First) //
	fmt.Println(p1.Last) //
	fmt.Println(p1.Age) // 0

	bs := []byte(`{"First": "James", "Last": "Bond", "wisdom score": 20}`)
	bs2 := []byte(`{"First": "Hope", "Last": "Kim", "Age": 25}`)

	json.Unmarshal(bs, &p1)
	json.Unmarshal(bs2, &p2)

	fmt.Println(" ------------------------------------------------- ")
	fmt.Println(p1.First, p1.Last, p1.Age) // James Bond 20
	fmt.Println(p2.First, p2.Last, p2.Age) // Hope Kim 0
}