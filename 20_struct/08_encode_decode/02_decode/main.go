package main

import (
	"strings"
	"encoding/json"
	"fmt"
)

type Person struct{
	First string
	Last string
	Age int
	notExported int
}

func main(){
	var p1 Person

	/*
	NewReader() takes a string and returns a pointer to a Reader type.
	*/
	rdr := strings.NewReader(`{"First": "James", "Last": "Bond", "Age": 20}`)

	/*
	NewDecoder() takes a Reader and returns a pointer to a Decoder.
	The encoding/json pkg provides a Decode method for Decoder pointers
	that stores its input into the pointer that it's passed.
	*/
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1)
}