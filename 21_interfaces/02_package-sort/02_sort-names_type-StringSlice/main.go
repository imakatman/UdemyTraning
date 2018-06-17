package main

import (
	"sort"
	"fmt"
)

func main(){
	studyGroup := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println(sort.StringsAreSorted(studyGroup))
}