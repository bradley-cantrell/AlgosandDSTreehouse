package main

func recursiveBinarySearch(list []int, target int) bool {
	if len(list) != 0 {
		midpoint := (len(list)) / 2
		if list[midpoint] == target {
			return true
		} else {
			if list[midpoint] < target {
				return recursiveBinarySearch(list[midpoint+1:], target)
			} else {
				if list[midpoint] > target {
					return recursiveBinarySearch(list[:midpoint], target)
				}
			}
		}
	}
	return false
}
