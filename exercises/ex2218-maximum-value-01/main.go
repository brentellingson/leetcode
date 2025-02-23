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
	var dp func(i, k int) int

	// dp means picking k elements from piles[i:]
	dp = func(i, k int) int {
		if k == 0 || i == len(piles) {
			return 0
		}
		res := dp(i+1, k)
		cur := 0
		for j := 0; j < len(piles[i]) && j < k; j++ {
			cur += piles[i][j]
			res = max(res, cur+dp(i+1, k-j-1))
		}
		return res
	}

	memoize := func(f func(i, k int) int) func(i, k int) int {
		memo := fill(len(piles)+1, k+1, -1)
		return func(i, k int) int {
			if memo[i][k] < 0 {
				memo[i][k] = f(i, k)
			}
			return memo[i][k]
		}
	}

	dp = memoize(dp)
	res := dp(0, k)
	return res
}

func fill(m, n, v int) [][]int {
	rval := make([][]int, m)
	for i := range m {
		rval[i] = make([]int, n)
		for j := range n {
			rval[i][j] = v
		}
	}
	return rval
}
