// package main implements the *minimum time to make rope colorful* exercise from leetcode
// see: https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
package main

import "fmt"

func main() {
	print("abaac", 1, 2, 3, 4, 5)
	print("abc", 1, 2, 3)
	print("aaa", 1, 2, 3)
	print("aabaa", 1, 2, 3, 4, 1)
}

func print(colors string, cost ...int) {
	fmt.Printf("%v, cost(%v) = %v\n", cost, colors, minCost(colors, cost))
}

func minCost(colors string, neededTime []int) int {
	rslt, max_cost := 0, 0
	for i := range []byte(colors) {
		if i > 0 && colors[i] != colors[i-1] {
			max_cost = 0
		}
		rslt += min(max_cost, neededTime[i])
		max_cost = max(max_cost, neededTime[i])
	}
	return rslt
}
