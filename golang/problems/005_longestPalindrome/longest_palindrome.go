package code

// ref https://leetcode.com/problems/longest-palindromic-substring/

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	longestLeftIndex := 0
	longestLength := 1
	for leftIndex := 0; leftIndex < length; leftIndex++ {
		for rightIndex := length - 1; rightIndex > leftIndex; rightIndex-- {
			checkLength := rightIndex - leftIndex + 1
			if checkLength <= longestLength {
				break
			}

			if isPalindrome(s, leftIndex, rightIndex) {
				longestLeftIndex = leftIndex
				longestLength = checkLength
			}
		}
	}

	return s[longestLeftIndex : longestLeftIndex+longestLength]
}

func isPalindrome(s string, leftIndex int, rightIndex int) bool {
	for {
		if s[leftIndex] != s[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
		if leftIndex >= rightIndex {
			break
		}
	}

	return true
}
