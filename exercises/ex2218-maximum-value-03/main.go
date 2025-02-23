// package main implements the *Maximum Value of K Coins From Piles* exercise from leetcode
// see: https://leetcode.com/problems/maximum-value-of-k-coins-from-piles
package main

import "fmt"

func main() {
	print([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7)
	print([][]int{{1, 100, 3}, {7, 8, 9}}, 2)
}

func print(piles [][]int, k int) {
	fmt.Printf("maxValueOfCoins(%v, %v) = %v\n", piles, k, maxValueOfCoins(piles, k))
}

func maxValueOfCoins(piles [][]int, k int) int {
	var memo [][]int
	for range piles {
		memo = append(memo, make([]int, k+1))
	}

	// dp(i, k) is result of allocating k coins to pile i, ... , len(piles)
	var dp func(i, k int) int
	dp = func(i, k int) int {
		if i >= len(piles) || k == 0 {
			return 0
		}
		if memo[i][k] > 0 {
			return memo[i][k]
		}
		rslt := dp(i+1, k)
		sum := 0
		for j := 0; j < k && j < len(piles[i]); j++ {
			sum += piles[i][j]
			rslt = max(rslt, sum+dp(i+1, k-j-1))
		}
		memo[i][k] = rslt
		return rslt
	}

	return dp(0, k)
}
