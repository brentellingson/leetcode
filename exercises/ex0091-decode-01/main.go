// package main implements the *decode ways* exercise from leetcode
// see: https://leetcode.com/problems/decode-ways
package main

import "fmt"

func main() {
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
	r := []rune(s)
	dp := make([]int, len(r))

	for i := range r {
		switch i {
		case 0:
			dp[i] = one(r[i], 1)
		case 1:
			dp[i] = one(r[i], dp[i-1]) + two(r[i-1], r[i], 1)
		default:
			dp[i] = one(r[i], dp[i-1]) + two(r[i-1], r[i], dp[i-2])
		}
		if dp[i] == 0 {
			return 0
		}
	}
	return dp[len(r)-1]
}

func one(r rune, val int) int {
	if '1' <= r && r <= '9' {
		return val
	}
	return 0
}

func two(p, r rune, val int) int {
	if p == '1' || (p == '2' && '0' <= r && r <= '6') {
		return val
	}
	return 0
}
