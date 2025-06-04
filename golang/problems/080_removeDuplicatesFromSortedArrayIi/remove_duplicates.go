package code

// ref https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

// Example 1:
// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]
// Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// Example 2:
// Input: nums = [0,0,1,1,1,1,2,3,3]
// Output: 7, nums = [0,0,1,1,2,3,3,_,_]
// Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
// It does not matter what you leave beyond the returned k (hence they are underscores).

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	k := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[k-2] && nums[i] == nums[i-1] {
			continue
		}

		nums[k] = nums[i]
		k++
	}

	return k
}
