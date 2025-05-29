package code

import (
	"math"
)

// ref https://leetcode.com/problems/reverse-integer/

// Example 1:

// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21

func reverse(x int) int {
	result := 0
	for integer := x; integer != 0; integer = integer / 10 {
		digit := integer % 10
		result = result*10 + digit
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
