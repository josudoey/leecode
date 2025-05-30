package code

import (
	"math"
	"regexp"
	"strconv"
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

var integerRegexp = regexp.MustCompile("^\\s*([+-]?[0-9]+).*")

func myAtoi(s string) int {
	result := 0

	matchedStrings := integerRegexp.FindStringSubmatch(s)
	if len(matchedStrings) == 0 {
		return result
	}

	n, _ := strconv.ParseInt(matchedStrings[1], 10, 64)
	if n > math.MaxInt32 {
		return math.MaxInt32
	} else if n < math.MinInt32 {
		return math.MinInt32
	}

	return int(n)
}
