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
	for _, num := range nums {
		numSet[num] = numSet[num] + 1
	}

	var maxNum, maxCount int

	for num, count := range numSet {
		if maxCount >= count {
			continue
		}

		maxCount = count
		maxNum = num
	}

	return maxNum
}
