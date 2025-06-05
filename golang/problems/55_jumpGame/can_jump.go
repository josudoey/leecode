package code

// ref https://leetcode.com/problems/jump-game/

// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

// Example 2:
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

func canJump(nums []int) bool {
	// only one; nums = [n]
	if len(nums) < 2 {
		return true
	}

	// first is zero; nums = [0,...]
	if nums[0] == 0 {
		return false
	}

	lastIndex := len(nums) - 1
	maxReachIndex := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > maxReachIndex {
			return false
		}

		currentReachIndex := i + nums[i]
		if currentReachIndex >= lastIndex {
			return true
		}

		if currentReachIndex > maxReachIndex {
			maxReachIndex = currentReachIndex
		}
	}

	return true
}
