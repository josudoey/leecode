package code

import (
	"math"
)

// ref https://leetcode.com/problems/string-to-integer-atoi/

// Example 1:
// Input: s = "42"
// Output: 42
// Explanation:
// The underlined characters are what is read in and the caret is the current reader position.
// Step 1: "42" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "42" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "42" ("42" is read in)
//            ^
// Example 2:
// Input: s = " -042"
// Output: -42
// Explanation:
// Step 1: "   -042" (leading whitespace is read and ignored)
//             ^
// Step 2: "   -042" ('-' is read, so the result should be negative)
//              ^
// Step 3: "   -042" ("042" is read in, leading zeros ignored in the result)
//                ^
// Example 3:
// Input: s = "1337c0d3"
// Output: 1337
// Explanation:
// Step 1: "1337c0d3" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "1337c0d3" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "1337c0d3" ("1337" is read in; reading stops because the next character is a non-digit)
//              ^
// Example 4:
// Input: s = "0-1"
// Output: 0
// Explanation:
// Step 1: "0-1" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "0-1" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "0-1" ("0" is read in; reading stops because the next character is a non-digit)
//           ^
// Example 5:
// Input: s = "words and 987"
// Output: 0

// Explanation:
// Reading stops at the first non-digit character 'w'.

func myAtoi(s string) int {
	currentIndex := 0
	for currentIndex < len(s) {
		if s[currentIndex] != ' ' {
			break
		}

		currentIndex++
	}

	if currentIndex == len(s) {
		return 0
	}

	result := 0
	sign := 1
	if s[currentIndex] == '-' {
		sign = -1
		currentIndex++
	} else if s[currentIndex] == '+' {
		currentIndex++
	}

	for currentIndex < len(s) {
		if s[currentIndex] < '0' || s[currentIndex] > '9' {
			break
		}

		digit := int(s[currentIndex] - '0')
		if result > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		currentIndex++
	}

	result *= sign

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}
