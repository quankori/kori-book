package sumtwonumber

import (
	"testing"
)

// Helper function to compare two linked lists for equality
func equal(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Value != l2.Value {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		l1 := createList(tt.l1)
		l2 := createList(tt.l2)
		expected := createList(tt.expected)
		result := AddTwoNumbers(l1, l2)

		if !equal(result, expected) {
			t.Errorf("AddTwoNumbers(%v, %v) = %v; want %v",
				tt.l1, tt.l2, listToSlice(result), tt.expected)
		}
	}
}

// Helper function to convert a linked list to a slice for easier comparison in error messages
func listToSlice(l *ListNode) []int {
	var result []int
	for l != nil {
		result = append(result, l.Value)
		l = l.Next
	}
	return result
}
