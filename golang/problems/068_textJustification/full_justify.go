package code

import (
	"strings"
)

// ref https://leetcode.com/problems/text-justification/

// Example 1:
// Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
// Output:
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]

// Example 2:
// Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
// Output:
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]
// Explanation: Note that the last line is "shall be    " instead of "shall     be", because the last line must be left-justified instead of fully-justified.
// Note that the second line is also left-justified because it contains only one word.

// Example 3:
// Input: words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"], maxWidth = 20
// Output:
// [
//   "Science  is  what we",
//   "understand      well",
//   "enough to explain to",
//   "a  computer.  Art is",
//   "everything  else  we",
//   "do                  "
// ]

func fullJustify(words []string, maxWidth int) []string {
	var result []string

	for len(words) > 0 {
		wordCount := len(words[0])
		i := 1

		for ; i < len(words); i++ {
			if wordCount+len(words[i])+i-1 >= maxWidth {
				break
			}

			wordCount += len(words[i])
		}

		if i != len(words) {
			result = append(result, concat(words[0:i], maxWidth-wordCount))
			words = words[i:]

			continue
		}

		lastLine := strings.Join(words[0:i], " ")
		result = append(result, lastLine+strings.Repeat(" ", maxWidth-len(lastLine)))
		break
	}

	return result
}

func concat(words []string, spaceWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaceWidth)
	}
	wordCount := len(words)
	avgSpace := spaceWidth / (wordCount - 1)
	remainingSpace := spaceWidth % (wordCount - 1)

	result := make([]string, len(words)*2-1)
	for i := 0; i < len(words)-1; i++ {
		result[i*2] = words[i]
		spaceCount := avgSpace
		if i < remainingSpace {
			spaceCount++
		}
		result[i*2+1] = strings.Repeat(" ", spaceCount)
	}
	result[2*(len(words)-1)] = words[len(words)-1]

	return strings.Join(result, "")
}
