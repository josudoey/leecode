package code

import (
	"sort"
)

// ref https://leetcode.com/problems/3sum/

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		first := nums[i]
		leftIndex, rightIndex := i+1, len(nums)-1
		for leftIndex < rightIndex {
			second, third := nums[leftIndex], nums[rightIndex]
			subtotal := second + third

			if -first != subtotal {
				if -first < subtotal {
					rightIndex--
				} else {
					leftIndex++
				}

				continue
			}

			result = append(result, []int{first, second, third})

			for leftIndex < rightIndex && second == nums[leftIndex+1] {
				leftIndex++
			}
			for leftIndex < rightIndex && third == nums[rightIndex-1] {
				rightIndex--
			}
			leftIndex++
			rightIndex--
		}
	}

	return result
}
