package main

import "fmt"

func main(){
	records := make([][]string, 0)

	fmt.Println(len(records))
	fmt.Println(cap(records))

	record1 := []string{
		"Hope Kim",
		"05/11/1993",
		"Los Angeles, CA",
	}

	record2 := []string{
		"Pearl Iwao",
		"04/06/1983",
		"Seoul, Korea",
	}

	records = append(records, record1)
	records = append(records, record2)

	fmt.Println(records)
	fmt.Println(len(records))
	fmt.Println(cap(records))
}