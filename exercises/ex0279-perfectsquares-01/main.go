// package main solves the *perfect squares* exercise from leetcode
// see: https://leetcode.com/problems/perfect-squares/
package main

import (
	"fmt"
)

func main() {
	for i := range 10_000 {
		fmt.Printf("numSquares(%v) = %v\n", i+1, numSquares(i+1))
	}
}

// numsquares implemented by sequential computation
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}
	return dp[n]
}
