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
		buckets[n]++
	}

	// The difference in order of magnitude is 3 now
	fmt.Println(buckets)
}

func hashTable(word string, buckets int) int{
	var sum int

	for _, v := range word{
		sum += int(v)
	}

	return sum % buckets
}