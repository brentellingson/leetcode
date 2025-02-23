// package main implements the *decode ways* exercise from leetcode
// see: https://leetcode.com/problems/decode-ways
package main

import "fmt"

func main() {
	print(1, 2)
	print(1, 2)
	print(1, 2, 3, 1)
	print(2, 7, 9, 3, 1)
	print(100, 1, 1, 100, 1, 1, 100)
}

func print(nums ...int) {
	fmt.Printf("rob(%v) = %v\n", nums, rob(nums))
}

// robt a house with iterative construction
func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		switch i {
		case 0:
			dp[i] = max(nums[i])
		case 1:
			dp[i] = max(nums[i], dp[i-1])
		default:
			dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		}
	}
	return dp[len(nums)-1]
}
