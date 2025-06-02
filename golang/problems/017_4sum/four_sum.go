package code

import (
	"sort"
)

// ref https://leetcode.com/problems/4sum/

// Example 1:
// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

// Example 2:
// Input: nums = [2,2,2,2,2], target = 8
// Output: [[2,2,2,2]]

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {
		first := nums[i]
		if i > 0 && first == nums[i-1] {
			continue
		}

		if first+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		if first+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			second := nums[j]
			if j > i+1 && second == nums[j-1] {
				continue
			}

			sum12 := first + second
			if sum12+nums[j+1]+nums[j+2] > target {
				break
			}

			if sum12+nums[n-2]+nums[n-1] < target {
				continue
			}

			remainingTarget := target - first - second

			leftIndex, rightIndex := j+1, n-1
			for leftIndex < rightIndex {
				third, fourth := nums[leftIndex], nums[rightIndex]
				sum34 := third + fourth

				if sum34 != remainingTarget {
					if sum34 < remainingTarget {
						leftIndex++
					} else {
						rightIndex--
					}
					continue
				}

				result = append(result, []int{first, second, third, fourth})

				for leftIndex < rightIndex && nums[leftIndex] == nums[leftIndex+1] {
					leftIndex++
				}
				for leftIndex < rightIndex && nums[rightIndex] == nums[rightIndex-1] {
					rightIndex--
				}

				leftIndex++
				rightIndex--
			}
		}
	}

	return result
}
