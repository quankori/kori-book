package arrays

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
*/
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	mArea := 0
	for left < right {
		var minHeight int
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}
		if mArea < (minHeight * (right - left)) {
			mArea = minHeight * (right - left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return mArea
}
