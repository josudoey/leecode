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
	palindromeCache := map[string]struct{}{}

	size := len(s)
	if size == 0 {
		return ""
	}

	longest := s[0:1]
	longestSize := 1
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			substrSize := j - i + 1
			if substrSize <= longestSize {
				continue
			}

			substr := s[i : j+1]
			if checkPalindrome(substr, palindromeCache) {
				palindromeCache[substr] = struct{}{}
				longest = substr
				longestSize = substrSize
			}
		}
	}

	return longest
}

func checkPalindrome(s string, palindromeCache map[string]struct{}) bool {
	size := len(s)
	if size < 2 {
		return size == 1
	}

	if _, ok := palindromeCache[s]; ok {
		return true
	}

	halfSize := size / 2
	leftIndex := halfSize - 1
	rightIndex := halfSize

	if size%2 == 1 {
		rightIndex += 1
	}

	for i := 0; i < halfSize; i++ {
		if s[leftIndex-i] != s[rightIndex+i] {
			return false
		}
	}

	return true
}
