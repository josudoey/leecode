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

func hash(key int) int {
	const A = 2654435769        // 2^32 / phi
	const primeMod = 4294967291 // 2^32 - 5

	hash := (uint64(key) * A) % uint64(primeMod)
	return int(hash)
}

func findSubstring(s string, words []string) []int {
	var (
		result = []int{}

		wordCountMap  = map[string]int{}
		wordVectorMap = map[string]int{}
	)

	uniqueCounter := 0
	for _, word := range words {
		if count, exists := wordCountMap[word]; exists {
			wordCountMap[word] = count + 1
			continue
		}

		uniqueCounter++
		wordCountMap[word] = 1
		wordVector := hash(uniqueCounter)
		wordVectorMap[word] = int(wordVector)
	}

	expectedCheckSum := 0
	for word, vector := range wordVectorMap {
		expectedCheckSum += (vector * wordCountMap[word])
	}

	lengthPerWord := len(words[0])
	substringLength := lengthPerWord * len(words)
	for i := 0; i <= len(s)-substringLength; i++ {
		checkSum := 0

		for j := 0; j < len(words); j++ {
			offset := i + (j * lengthPerWord)
			word := s[offset : offset+lengthPerWord]

			wordVector, ok := wordVectorMap[word]
			if !ok {
				break
			}
			checkSum += wordVector
		}

		if checkSum != expectedCheckSum {
			continue
		}

		result = append(result, i)
	}

	return result
}
