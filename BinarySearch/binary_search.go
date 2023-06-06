package main

import "fmt"

//binary search: worst case scenario O(log n) where n is the size of the array
//best case scenario O(1)
func main() {
	// list has to be sorted for binary search
	list := []int{1, 2, 4, 10, 19, 32}
	target := 4
	result := binarySearch(list, target)
	fmt.Println(result)
}

func binarySearch(list []int, target int) int {
	first := 0
	last := len(list) - 1

	for first <= last {
		middle := first + (last-first)/2
		if list[middle] == target {
			fmt.Println(middle)
			return middle
		} else if list[middle] < target {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}
	fmt.Println("target not found in list")
	return -1
}
