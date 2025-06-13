package code

// ref https://leetcode.com/problems/reverse-words-in-a-string/

// Example 1:
// Input: s = "the sky is blue"
// Output: "blue is sky the"

// Example 2:
// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.

// Example 3:
// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

func reverseWords(s string) string {
	var (
		n     = len(s)
		words [][]byte
	)

	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}

		headIndex := i
		for ; i < n; i++ {
			if s[i] != ' ' {
				continue
			}

			break
		}

		words = append(words, []byte(s[headIndex:i]))
	}

	result := []byte{}
	for len(words) > 0 {
		result = append(result, words[len(words)-1]...)
		result = append(result, ' ')
		words = words[:len(words)-1]
	}

	return string(result[:len(result)-1])
}
