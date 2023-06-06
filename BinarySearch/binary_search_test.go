package main

import "testing"

func Test_binarySearch(t *testing.T) {
	list := []int{1, 2, 4, 10, 19, 32}

	//Case: target is in position 0
	target := 1
	expected := 0
	result := binarySearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}

	//Case: target is in position 6
	target = 32
	expected = 5
	result = binarySearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}

	//Case: target is in position 2
	target = 4
	expected = 2
	result = binarySearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}

	//Case: target is not present
	target = 22
	expected = -1
	result = binarySearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}

	//Case: empty list
	emptyList := []int{}
	target = 1
	expected = -1
	result = binarySearch(emptyList, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}
