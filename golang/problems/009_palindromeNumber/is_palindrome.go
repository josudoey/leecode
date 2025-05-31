package code

import "math"

// ref https://leetcode.com/problems/palindrome-number/

// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	log10x := math.Log10(float64(x))
	leftDigitDivisor := int(math.Pow10(int(log10x)))
	for x != 0 {
		if leftDigitDivisor == 1 {
			break
		}

		rightDigit := x % 10
		leftDigit := x / leftDigitDivisor

		if rightDigit != leftDigit {
			return false
		}

		x = (x % leftDigitDivisor) / 10
		leftDigitDivisor = leftDigitDivisor / 100
	}

	return true
}
