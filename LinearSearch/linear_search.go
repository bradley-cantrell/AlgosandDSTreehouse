package main

import (
	"fmt"
)

// linear search: worst case scenario = O(n) because in the worst case every value has to be checked and the value may not be present at all
// best case scenario = O(1) if the value is the first value in the slice

func main() {
	list := []int{1, 2, 4, 10, 19, 7, 32}
	target := 4
	linearSearch(list, target)
}

func linearSearch(list []int, target int) int {
	for i, value := range list {
		if value == target {
			fmt.Println(i)
			return i
		}
	}
	fmt.Println("index not present")
	return -1
}
