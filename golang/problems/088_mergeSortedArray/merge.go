package code

// ref https://leetcode.com/problems/merge-sorted-array/

// Example 1:
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

// Example 2:
// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].

// Example 3:
// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The result of the merge is [1].
// Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

func merge(nums1 []int, m int, nums2 []int, n int) {
	index, index1, index2 := m+n-1, m-1, n-1
	for index >= 0 && index2 >= 0 {
		if index1 > 0 && nums1[index1] > nums2[index2] {
			nums1[index] = nums1[index1]
			index1 -= 1
		} else {
			nums1[index] = nums2[index2]
			index2 -= 1
		}

		index--
	}
}
