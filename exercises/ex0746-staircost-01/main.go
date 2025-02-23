// package main implements the *min cost climbing stairs* exercise from leetcode
// see: https://leetcode.com/problems/min-cost-climbing-stairs/
package main

import "fmt"

func main() {
	print(1, 2)
	print(1, 1, 1)
	print(1, 1, 1, 1)
	print(1, 100, 1)
	print(10, 15, 20)
	print(1, 100, 1, 1, 1, 100, 1, 1, 100, 1)
}

func print(cost ...int) {
	fmt.Printf("cost(%v) = %v\n", cost, minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		a, b = b, cost[i]+min(a, b)
	}
	return min(a, b)
}
