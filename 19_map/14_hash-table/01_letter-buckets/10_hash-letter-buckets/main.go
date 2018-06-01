package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main() {
	//resp, err := http.Get("https://www.gutenberg.org/files/2554/2554-0.txt")
	resp, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	// If there is no error in the Get request, then the value is nil
	if err != nil {
		log.Fatal(err)
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}

	/*
	 Create a Scanner, which is an interface for reading data, which takes a Reader interface
	 as an argument. resp.Body is a ReadCloser interface which means its both a Reader interface
	 and a Closer interface. Shorthand for interface is io
	 */
	scanner := bufio.NewScanner(resp.Body)
	// When you get a response, you need to make sure to close it before the function ends.
	defer resp.Body.Close()

	// Split Crime and Punishment into words
	scanner.Split(bufio.ScanWords)

	// Create buckets for our words, which will be a slice of length and capacity 200,
	// of type int which will be our buckets.
	buckets := make([]int, 200)

	// Loop over the words
	for scanner.Scan(){
		n := hashBucket(scanner.Text())
		buckets[n]++
		/*
		 Grab the UTF-8 number of the first letter of each word,
		 add one because an array is 0-indexed, and then create
		 a bucket at that index in the buckets.
		*/
	}

	// Print slices from A - Z
	fmt.Println(buckets[65:123])
	/*
	  Output: [1790 1260 910 522 325 658 473 812 2829 244 80 504 563 672 433 816 291 259 1671 1771 87 119 1080 5 205 15 21 0 0 0 2 0 21787 9692 7385 5219 3614 7576 3016 12620 11607 582 852 5321 7765 4060 13798 5256 395 3592 16260 33335 2591 1452 13314 1 2272 19]

	  This means that there are 1790 words that start with capital letter "A", 1260
	  letters that start with capital letter "B", etc.. etc...

	  Purpose of this is to figure out how many entries are going to be within each
	  bucket. This is a very uneven distribution.
	*/
}

func hashBucket(word string) int{
	// Convert the first letter of the word into a "Decimal" number and return that
	return int(word[0])
}