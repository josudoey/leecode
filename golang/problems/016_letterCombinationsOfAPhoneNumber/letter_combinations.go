package code

import "strings"

// ref https://leetcode.com/problems/letter-combinations-of-a-phone-number/

// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Example 2:
// Input: digits = ""
// Output: []

// Example 3:
// Input: digits = "2"
// Output: ["a","b","c"]

var digitLetterMap = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := strings.Split(digitLetterMap[digits[0]], "")
	for _, digit := range digits[1:] {
		var combined []string
		for _, prefix := range result {
			for _, letter := range digitLetterMap[uint8(digit)] {
				combined = append(combined, prefix+string(letter))
			}
		}

		result = combined
	}

	return result
}
