package code

import "sort"

// ref https://leetcode.com/problems/3sum-closest/

// Example 1:
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

// Example 2:
// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	closestDelta := abs(target - closestSum)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		first := nums[i]
		leftIndex, rightIndex := i+1, len(nums)-1
		for leftIndex < rightIndex {
			second, third := nums[leftIndex], nums[rightIndex]
			sum := first + second + third

			if target == sum {
				return target
			}

			delta := abs(target - sum)
			if closestDelta > delta {
				closestSum = sum
				closestDelta = delta
			}

			if target < sum {
				rightIndex--
			} else {
				leftIndex++
			}

		}
	}

	return closestSum
}
