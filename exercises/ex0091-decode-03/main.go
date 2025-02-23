// package main implements the *decode ways* exercise from leetcode
// see: https://leetcode.com/problems/decode-ways
package main

import "fmt"

func main() {
	print("0")
	print("1")
	print("2")
	print("3")
	print("12")
	print("226")
	print("06")
	print("911")
	print("26")
	print("27")
	print("262626")
}

func print(s string) {
	fmt.Printf("numDeconds(%v) = %v\n", s, numDecodings(s))
}

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	for i := len(s); i >= 0; i-- {
		s := s[i:]
		if len(s) == 0 {
			dp[i] = 1
		} else if s[0] == '0' {
			dp[i] = 0
		} else if len(s) > 1 && s[0] == '1' {
			dp[i] = dp[i+1] + dp[i+2]
		} else if len(s) > 1 && s[0] == '2' && s[1] <= '6' {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}
