package code

// ref https://leetcode.com/problems/substring-with-concatenation-of-all-words/

// Example 1:
// Input: s = "barfoothefoobarman", words = ["foo","bar"]
// Output: [0,9]
// Explanation:
// The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
// The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

// Example 2:
// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// Output: []
// Explanation:
// There is no concatenated substring.

// Example 3:
// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
// Output: [6,9,12]
// Explanation:
// The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
// The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
// The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].

func findSubstring(s string, words []string) []int {
	var (
		result = []int{}

		wordCountMap = map[string]int{}
	)

	for _, word := range words {
		wordCountMap[word]++
	}

	lengthPerWord := len(words[0])
	substringLength := lengthPerWord * len(words)
	for i := 0; i <= len(s)-substringLength; i++ {
		currentWordCount := map[string]int{}

		for j := 0; j < len(words); j++ {
			offset := i + (j * lengthPerWord)
			word := s[offset : offset+lengthPerWord]

			count, ok := wordCountMap[word]
			if !ok {
				break
			}

			currentWordCount[word]++
			if currentWordCount[word] > count {
				break
			}
		}

		if len(currentWordCount) != len(wordCountMap) {
			continue
		}

		matched := true
		for word, count := range wordCountMap {
			if currentWordCount[word] == count {
				continue
			}

			matched = false
			break
		}

		if !matched {
			continue
		}

		result = append(result, i)
	}

	return result
}
