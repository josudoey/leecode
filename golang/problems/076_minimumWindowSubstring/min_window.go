package code

// ref https://leetcode.com/problems/minimum-window-substring/

// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.

// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.

func minWindow(s string, t string) string {
	var (
		m, n            = len(s), len(t)
		minWindowLength = len(s) + 1
		minWindowIndex  = -1
		charCounts      = make([]int, 52)
		targetCharCount = 0
	)

	for i := 0; i < n; i++ {
		charCode := compressChar(t[i])
		if charCounts[charCode] == 0 {
			targetCharCount++
		}

		charCounts[charCode]++
	}

	leftIndex := 0
	windowCharCounts := make([]int, 52)
	matchedCharCount := 0
	for rightIndex := 0; rightIndex < m; rightIndex++ {
		rightCharCode := compressChar(s[rightIndex])
		if rightCharCode < 0 {
			continue
		}

		windowCharCounts[rightCharCode]++
		if windowCharCounts[rightCharCode] == charCounts[rightCharCode] {
			matchedCharCount++
		}

		if matchedCharCount != targetCharCount {
			continue
		}

		for ; leftIndex <= rightIndex; leftIndex++ {
			leftCharCode := compressChar(s[leftIndex])
			if leftIndex < 0 {
				continue
			}

			windowCharCounts[leftCharCode]--
			if windowCharCounts[leftCharCode] >= charCounts[leftCharCode] {
				continue
			}

			matchedCharCount--
			windowLength := rightIndex - leftIndex + 1
			if windowLength > minWindowLength {
				leftIndex++
				break
			}

			minWindowLength = windowLength
			minWindowIndex = leftIndex
			leftIndex++
			break
		}

	}

	if minWindowLength > m {
		return ""
	}

	return s[minWindowIndex : minWindowIndex+minWindowLength]
}

func compressChar(char byte) int {
	if 'a' <= char && char <= 'z' {
		return int(char - 'a' + 26)
	}

	if 'A' <= char && char <= 'Z' {
		return int(char - 'A')
	}
	return -1
}
