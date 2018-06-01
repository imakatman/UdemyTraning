package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main(){
	resp, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil{
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(resp.Body)

	defer resp.Body.Close()

	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 12)

	for scanner.Scan(){
		n := hashTable(scanner.Text(), 12)
		//fmt.Println(n)
		buckets[n]++
	}

	// The difference in order of magnitude is 10 which is still pretty high
	fmt.Println(buckets)
}

func hashTable(word string, buckets int) int{
	letter := int(word[0])
	bucket := letter % buckets

	return bucket
}