package main

import (
	"sort"
	"fmt"
)

type numbers []int

// Attaching method Len() to type numbers
func (n numbers) Len() int{ return len(n)}

// Attaching method Less() to type numbers
func (n numbers) Less(i, j int) bool{ return n[i] < n[j]}

// Attaching method Swap() to type numbers
func (n numbers) Swap(i, j int) {n[i], n[j] = n[j], n[i]}

func main(){
	n := numbers{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(sort.Reverse(n))
	fmt.Println(n)
}