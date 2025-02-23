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

// maxValueOfCoins computed by memoization
//  1. first compute sums[i][j] := total value of first j coins in pile i
//  2. next compute  memo[i][j] := max value of selecting j coins from piles i...len(piles)-1
//     loop backwards through i, and
//     determine memo[i][j] := max(mem[i+1][n] + sum[i][j-n]) for 0 <= n <= k
//  3. at the end, look at mem[0][k] and return it
func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	// compute sum[i][j] := total value of first j coins in pile i
	sums := make([][]int, n)
	for i := 0; i < n; i++ {
		// add one to number of coins in pile because we can select 0 coins
		sums[i] = make([]int, len(piles[i])+1)
		sums[i][0] = 0
		for j := 0; j < len(piles[i]); j++ {
			sums[i][j+1] = sums[i][j] + piles[i][j]
		}
		fmt.Printf("sums(%v) = %v\n", piles[i], sums[i])
	}

	// compute memo[i][j] := max value of selecting j coins from pile i...n
	// we add an empty memo on the end to make computation simpler
	memo := make([][]int, n+1)
	memo[n] = []int{0}
	for i := n - 1; i >= 0; i-- {
		size := min(k+1, len(sums[i])+len(memo[i+1])-1)
		memo[i] = make([]int, size)
		for j := 0; j < size; j++ {
			memo[i][j] = 0
			for n := 0; n <= j; n++ {
				// we could put this in the loop boundries, but the math sucks
				if j-n >= 0 && j-n < len(sums[i]) && n < len(memo[i+1]) {
					memo[i][j] = max(memo[i][j], sums[i][j-n]+memo[i+1][n])
				}
			}
		}
		fmt.Printf("memo(%v) = %v\n", i, memo[i])
	}

	return memo[0][k]
}
