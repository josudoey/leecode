package lengthOfLongestSubstring

// ref https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	var (
		result       int
		charIndexMap = map[rune]int{}
	)

	for index, char := range s {
		dupCharIndex, ok := charIndexMap[char]
		if !ok {
			charIndexMap[char] = index
			continue
		}

		currentLen := len(charIndexMap)
		if currentLen > result {
			result = len(charIndexMap)
		}
		removePrevSubstring(charIndexMap, dupCharIndex)
		charIndexMap[char] = index

	}

	lastLength := len(charIndexMap)
	if lastLength > result {
		return lastLength
	}

	return result
}

func removePrevSubstring(charIndexMap map[rune]int, headIndex int) {
	for char, index := range charIndexMap {
		if index > headIndex {
			continue
		}
		delete(charIndexMap, char)
	}
}
