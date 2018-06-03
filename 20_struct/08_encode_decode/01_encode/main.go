package main

import (
	"encoding/json"
	"os"
)

type Person struct{
	First string
	Last string
	Age int
}

func main(){
	p1 := Person{"James", "Bond", 20}

	/*
	FROM DOCS... https://golang.org/pkg/encoding/json/#Encoder
	-------------------------------------------------------------
	NewEncoder() is a func that takes a Writer type and
	returns a pointer to an Encoder, *Encoder. If you have
	a pointer to an Encoder, then you can use the Encode method.

	Stdout is an open file that we can write to. It gives us NewFile()
	which returns a pointer to a file. You can then check the docs to see
	if the pointer has the Write method which is provided by the Writer
	interface. If it does, then it is implicitly a Writer interface.

	That's why you can stick the os.Stdout into the NewEncoder method,
	because its a Writer interface. This is polymorphism.

	You can use this "discovery" strategy by recursively going through the docs.
	*/
	json.NewEncoder(os.Stdout).Encode(p1); // {"First":"James","Last":"Bond","Age":20}
}