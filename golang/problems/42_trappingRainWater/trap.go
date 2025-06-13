package code

// ref https://leetcode.com/problems/trapping-rain-water/

// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

// Example 2:
// Input: height = [4,2,0,3,2,5]
// Output: 9

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	leftIndex := 0
	rightIndex := n - 1
	leftMax := height[leftIndex]
	rightMax := height[rightIndex]
	result := 0

	for leftIndex < rightIndex {
		if leftMax < rightMax {
			leftIndex++
			if height[leftIndex] > leftMax {
				leftMax = height[leftIndex]
				continue
			}

			result += leftMax - height[leftIndex]
			continue
		}

		rightIndex--
		if height[rightIndex] > rightMax {
			rightMax = height[rightIndex]
			continue
		}

		result += rightMax - height[rightIndex]
		continue
	}

	return result
}
