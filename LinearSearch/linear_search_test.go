package main

import "testing"

func Test_linearSearch(t *testing.T) {
	list := []int{1, 2, 4, 10, 19, 7, 32}

	//Case: target present at index 0
	target := 1
	expected := 0
	result := linearSearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}

	//Case: target present at index 6
	target = 32
	expected = 6
	result = linearSearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}

	//Case: target present at index 2
	target = 4
	expected = 2
	result = linearSearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}

	//Case: target not present
	target = 30
	expected = -1
	result = linearSearch(list, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
	
	//Case: empty list
	emptyList := []int{}
	target = 1
	expected = -1
	result = linearSearch(emptyList, target)
	if result != expected {
		t.Errorf("Expected index %d, but got %d", expected, result)
	}
}
