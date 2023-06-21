
package main

import "testing"

func Test_recursiveBinarySearch(t *testing.T) {
	list := []int{1, 2, 4, 10, 19, 32}
	
//Case: value is present, first position
target := 1
expected := true
result := recursiveBinarySearch(list, target)
if result != expected {
	t.Error("Expected:", expected, "Result:", result)
}

//Case: value is present, last position
target = 32
expected = true
result = recursiveBinarySearch(list, target)
if result != expected {
	t.Error("Expected:", expected, "Result:", result)
}

//Case: value is present, "random" position
target = 4
expected = true
result = recursiveBinarySearch(list, target)
if result != expected {
	t.Error("Expected:", expected, "Result:", result)
}

//Case: value is not present
target = 22
expected = false
result = recursiveBinarySearch(list, target)
if result != expected {
	t.Error("Expected:", expected, "Result:", result)
}

//Case: empty list
emptyList := []int{}
target = 1
expected = false
result = recursiveBinarySearch(emptyList, target)
if result != expected {
	t.Error("Expected:", expected, "Result:", result)
}

}