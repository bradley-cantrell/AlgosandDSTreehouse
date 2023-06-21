package main

import "fmt"

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
