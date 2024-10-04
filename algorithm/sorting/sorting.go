package sorting

func quickSort(arrays []int) []int {
	if len(arrays) < 2 {
		return arrays
	}
	left, right := 0, len(arrays)-1

	pivot := len(arrays) / 2

	arrays[pivot], arrays[right] = arrays[right], arrays[pivot]

	for i := range arrays {
		if arrays[i] < arrays[right] {
			arrays[i], arrays[left] = arrays[left], arrays[i]
			left++
		}
	}
	arrays[left], arrays[right] = arrays[right], arrays[left]
	quickSort(arrays[:left])
	quickSort(arrays[left+1:])
	return arrays
}
