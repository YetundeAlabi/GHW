package main

import (
	"fmt"
	"sort"
)

func main() {
	numSlice := []int{23, 1, 4, 7, 10, 15}
	sort.Ints(numSlice)
	fmt.Println(numSlice) //[1 4 7 10 15 23]
}