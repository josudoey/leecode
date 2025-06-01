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

type TripletFilter struct {
	set map[int]map[int]map[int]struct{}
}

func newTripletFilter() *TripletFilter {
	return &TripletFilter{
		set: map[int]map[int]map[int]struct{}{},
	}
}

func (f *TripletFilter) Test(n0 int, n1 int, n2 int) bool {
	m0, ok := f.set[n0]
	if !ok {
		return false
	}

	m1, ok := m0[n1]
	if !ok {
		return false
	}

	_, ok = m1[n2]
	return ok
}

func (f *TripletFilter) Add(n0 int, n1 int, n2 int) {
	m0, ok := f.set[n0]
	if !ok {
		m0 = map[int]map[int]struct{}{}
		f.set[n0] = m0
	}

	m1, ok := m0[n1]
	if !ok {
		m1 = map[int]struct{}{}
		m0[n1] = m1
	}

	m1[n2] = struct{}{}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	tripletFilter := newTripletFilter()

	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]
		leftIndex, rightIndex := i+1, len(nums)-1
		for leftIndex < rightIndex {
			second, third := nums[leftIndex], nums[rightIndex]
			subtotal := second + third

			if -first < subtotal {
				rightIndex--
			} else {
				leftIndex++
			}

			if -first != subtotal {
				continue
			}

			if tripletFilter.Test(first, second, third) {
				continue
			}

			tripletFilter.Add(first, second, third)
			result = append(result, []int{first, second, third})
		}
	}

	return result
}
