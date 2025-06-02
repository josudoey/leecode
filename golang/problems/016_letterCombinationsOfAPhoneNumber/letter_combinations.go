package code

// ref https://leetcode.com/problems/letter-combinations-of-a-phone-number/

// Example 1:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Example 2:
// Input: digits = ""
// Output: []

// Example 3:
// Input: digits = "2"
// Output: ["a","b","c"]

var digitsLetters = [][]string{
	{"a", "b", "c"},      // 2
	{"d", "e", "f"},      // 3
	{"g", "h", "i"},      // 4
	{"j", "k", "l"},      // 5
	{"m", "n", "o"},      // 6
	{"p", "q", "r", "s"}, // 7
	{"t", "u", "v"},      // 8
	{"w", "x", "y", "z"}, // 9
}

func newStringChan(items []string) <-chan string {
	result := make(chan string)
	go func() {
		for _, s := range items {
			result <- s
		}

		close(result)
	}()

	return result
}

func newStringSlice(in <-chan string) []string {
	var result []string
	for s := range in {
		result = append(result, s)
	}

	return result
}

func combin(prefix string, in <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for s := range in {
			result <- prefix + s
		}
		close(result)
	}()

	return result
}

func merge(in []<-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for _, c := range in {
			for s := range c {
				result <- s
			}
		}
		close(result)
	}()

	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	letters := make([][]string, len(digits))
	for i, c := range digits {
		index := int(c - '2')
		letters[i] = digitsLetters[index]
	}

	result := newStringChan(letters[0])
	for i := 1; i < len(letters); i++ {
		combiners := []<-chan string{}
		for prefix := range result {
			combiners = append(combiners, combin(prefix, newStringChan(letters[i])))
		}

		result = merge(combiners)
	}

	return newStringSlice(result)
}
