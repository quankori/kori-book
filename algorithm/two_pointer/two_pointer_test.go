package arrays

import (
	"testing"
)

func Test1(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea := maxArea(height)
	if maxArea != 49 {
		t.Errorf("Output should be 49 instead of %v", maxArea)
	}
}

func Test2(t *testing.T) {
	height := []int{1, 1}
	maxArea := maxArea(height)
	if maxArea != 1 {
		t.Errorf("Output should be 1 instead of %v", maxArea)
	}
}

func Test3(t *testing.T) {
	height := []int{1, 2, 1}
	maxArea := maxArea(height)
	if maxArea != 2 {
		t.Errorf("Output should be 2 instead of %v", maxArea)
	}
}
