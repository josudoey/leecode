package code

// ref https://leetcode.com/problems/jump-game-ii/

// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Example 2:
// Input: nums = [2,3,0,1,4]
// Output: 2

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	lastIndex := n - 1
	step := 0
	currentReachIndex := 0
	nextReachIndex := 0

	for i := 0; i < lastIndex; i++ {
		reachIndex := i + nums[i]
		if nextReachIndex < reachIndex {
			nextReachIndex = reachIndex
		}

		if i != currentReachIndex {
			continue
		}

		step++
		currentReachIndex = nextReachIndex
		if currentReachIndex >= lastIndex {
			return step
		}
	}

	return step
}
