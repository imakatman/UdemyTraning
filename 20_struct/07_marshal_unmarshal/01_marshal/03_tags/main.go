package main

import (
	"fmt"
	"encoding/json"
)

type Person struct{
	First string
	Last string `json:"-"` // this omits the attribute (but not the value of the attribute if its assigned)
	Age int `json:"wisdom age"` // renames Age to wisdomm age
}

func main(){
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs)) //{"First":"James","wisdom age":20}
}