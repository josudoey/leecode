package code

// ref https://leetcode.com/problems/container-with-most-water/

// Example 1:
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

// Example 2:
// Input: height = [1,1]
// Output: 1

func maxArea(height []int) int {
	var result int
	leftIndex, rightIndex := 0, len(height)-1

	for leftIndex < rightIndex {

		h := min(height[leftIndex], height[rightIndex])
		w := rightIndex - leftIndex
		area := w * h
		result = max(result, area)

		if height[leftIndex] > height[rightIndex] {
			rightIndex--
		} else {
			leftIndex++
		}
	}

	return result
}
