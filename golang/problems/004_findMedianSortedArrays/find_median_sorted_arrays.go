package code

import (
	"sort"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	all := append(nums1, nums2...)
	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})

	size := len(all)
	medianIndex := size / 2
	if size%2 == 1 {
		return float64(all[medianIndex])
	}
	return float64(all[medianIndex-1]+all[medianIndex]) / 2
}
