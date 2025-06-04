package code

// ref https://leetcode.com/problems/majority-element/

// Example 1:
// Input: nums = [3,2,3]
// Output: 3

// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

func majorityElement(nums []int) int {
	numSet := map[int]int{}
	majorityCount := len(nums) / 2

	for _, n := range nums {
		numSet[n] = numSet[n] + 1
		if numSet[n] > majorityCount {
			return n
		}
	}

	return 0
}
