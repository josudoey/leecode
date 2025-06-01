package code

import "strings"

// ref https://leetcode.com/problems/longest-common-prefix/

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {
	var result strings.Builder
	if len(strs) == 0 {
		return result.String()
	}

	minLength := len(strs[0])
	for _, str := range strs[1:] {
		minLength = min(minLength, len(str))
	}

	for i := 0; i < minLength; i++ {
		char := strs[0][i]
		for _, str := range strs[1:] {
			if str[i] == char {
				continue
			}
			return result.String()
		}

		result.WriteByte(char)
	}

	return result.String()
}
