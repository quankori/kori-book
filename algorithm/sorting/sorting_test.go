package sorting

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	arraySorted := quickSort(array)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 8}
	if !reflect.DeepEqual(arraySorted, expected) {
		t.Errorf("Output should be %v instead of %v", expected, arraySorted)
	}
}
