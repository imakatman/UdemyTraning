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

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)

	for i := 0; i < 12; i++{
		buckets = append(buckets, []string{})
	}

	for scanner.Scan(){
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	for i := 0; i < 12; i++{
		fmt.Println(i, "len:", len(buckets[i]))
	}
}

func hashBucket(word string, buckets int) int{
	var sum int
	for _, v := range word{
		sum += int(v)
	}

	return sum % buckets
}