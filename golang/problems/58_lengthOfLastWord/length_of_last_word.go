package code

// ref https://leetcode.com/problems/length-of-last-word/

// Example 1:
// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.

// Example 2:
// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.

// Example 3:
// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.

func lengthOfLastWord(s string) int {
	n := len(s)
	result := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if result > 0 {
				return result
			}
			continue
		}

		result++
		if i == 0 {
			return result
		}

	}

	return result
}
