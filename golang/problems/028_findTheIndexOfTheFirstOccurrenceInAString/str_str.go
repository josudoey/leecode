package code

// ref https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

// Example 1:
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.

// Example 2:
// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.

func strStr(haystack string, needle string) int {
	n := len(haystack)
	for i := 0; i < n-len(needle)+1; i++ {
		if haystack[i] != needle[0] {
			continue
		}

		headIndex := i
		matchedCount := 1
		for ; matchedCount < len(needle); matchedCount++ {
			if haystack[i+matchedCount] != needle[matchedCount] {
				break
			}
		}

		if matchedCount != len(needle) {
			continue
		}

		return headIndex
	}
	return -1
}
