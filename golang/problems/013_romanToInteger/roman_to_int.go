package code

// ref https://leetcode.com/problems/roman-to-integer/

// Example 1:
// Input: s = "III"
// Output: 3
// Explanation: III = 3.

// Example 2:
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.

// Example 3:
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

var romanNumeralMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0
	for index := 0; index < len(s)-1; index++ {
		value := romanNumeralMap[s[index]]
		nextIndex := index + 1
		nextValue := romanNumeralMap[s[nextIndex]]
		if nextValue > value {
			result -= value
			continue
		}

		result += value
	}

	result += romanNumeralMap[s[len(s)-1]]
	return result
}
