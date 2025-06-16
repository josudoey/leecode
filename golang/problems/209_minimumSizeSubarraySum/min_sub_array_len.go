package code

// ref https://leetcode.com/problems/minimum-size-subarray-sum/

// Example 1:
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.

// Example 2:
// Input: target = 4, nums = [1,4,4]
// Output: 1

// Example 3:
// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)

	sum := 0
	leftIndex := 0
	minSubarrayLength := n + 1

	for rightIndex := 0; rightIndex < n; rightIndex++ {
		sum += nums[rightIndex]

		for sum >= target {
			currentLen := rightIndex - leftIndex + 1
			if currentLen == 1 {
				return 1
			}

			minSubarrayLength = min(currentLen, minSubarrayLength)
			sum -= nums[leftIndex]
			leftIndex++
		}
	}

	if minSubarrayLength <= n {
		return minSubarrayLength
	}

	return 0
}
