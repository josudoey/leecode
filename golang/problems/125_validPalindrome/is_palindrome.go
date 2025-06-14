package code

// ref https://leetcode.com/problems/valid-palindrome/

// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

func isPalindrome(s string) bool {
	leftIndex := 0
	rightIndex := len(s) - 1

	for leftIndex < rightIndex {
		leftChar := toLower(s[leftIndex])
		if !isAlphanumeric(leftChar) {
			leftIndex++
			continue
		}

		rightChar := toLower(s[rightIndex])
		if !isAlphanumeric(rightChar) {
			rightIndex--
			continue
		}

		if leftChar != rightChar {
			return false
		}

		leftIndex++
		rightIndex--
	}

	return true
}

func isAlphanumeric(c byte) bool {
	if ('a' <= c && c <= 'z') || '0' <= c && c <= '9' {
		return true
	}

	return false
}

func toLower(c byte) byte {
	if 'A' <= c && c <= 'Z' {
		c += 'a' - 'A'
		return c
	}

	return c
}
