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
	result := nums1
	mSlice, nSlice := make([]int, len(nums1)), nums2
	copy(mSlice, nums1)

	index, index1, index2 := 0, 0, 0
	for index1 < m && index2 < n {
		if mSlice[index1] <= nSlice[index2] {
			result[index] = mSlice[index1]
			index++
			index1++
		} else {
			result[index] = nSlice[index2]
			index++
			index2++
		}
	}

	for index1 < m {
		result[index] = mSlice[index1]
		index++
		index1++
	}

	for index2 < n {
		result[index] = nSlice[index2]
		index++
		index2++
	}
}
