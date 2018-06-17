package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	/*
	http's Get method returns a pointer to a Response type and an error, if there is one.
	*/
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil{
		log.Fatalln(err)
	}
	/*
	The Response type comes with many fields, one of them is Body
	which is a ReadCloser interface type. The ReadCloser interface
	type is an embedded type of both Reader and Writer.

 	ioutil's ReadAll method takes a Reader type and returns a slice
 	of bytes and an error if there is one.
	*/
	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)
}