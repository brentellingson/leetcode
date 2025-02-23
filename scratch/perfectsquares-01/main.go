// package main solves the *perfect squares* exercise from leetcode
// see: https://leetcode.com/problems/perfect-squares/
package main

import (
	"fmt"
)

func main() {
	for i, v := range squares(10) {
		fmt.Printf("squares(%v) = %v\n", i, v)
	}
}

// numsquares implemented by sequential computation
func squares(n int) [][]int {
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = append(dp[i-1], 1)
		for j := 2; j*j <= i; j++ {
			if len(dp[i-j*j])+1 < len(dp[i]) {
				dp[i] = append(dp[i-j*j], j*j)
			}
		}
	}
	return dp
}
