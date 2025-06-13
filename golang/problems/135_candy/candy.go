package code

// ref https://leetcode.com/problems/candy/

// Example 1:
// Input: ratings = [1,0,2]
// Output: 5
// Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

// Example 2:
// Input: ratings = [1,2,2]
// Output: 4
// Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
// The third child gets 1 candy because it satisfies the above two conditions.

func candy(ratings []int) int {
	n := len(ratings)
	candyCounts := make([]int, n)
	for i := 0; i < n; i++ {
		candyCounts[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candyCounts[i] = candyCounts[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyCounts[i] = max(candyCounts[i], candyCounts[i+1]+1)
		}
	}

	result := 0
	for _, count := range candyCounts {
		result += count
	}

	return result
}
