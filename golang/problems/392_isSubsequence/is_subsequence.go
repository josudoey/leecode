package code

// ref https://leetcode.com/problems/is-subsequence/

// Example 1:
// Input: s = "abc", t = "ahbgdc"
// Output: true

// Example 2:
// Input: s = "axc", t = "ahbgdc"
// Output: false

func isSubsequence(s string, t string) bool {
	indexA := 0
	for indexB := 0; indexA < len(s) && indexB < len(t); indexB++ {
		if s[indexA] != t[indexB] {
			continue
		}
		indexA++
	}

	return indexA == len(s)
}
