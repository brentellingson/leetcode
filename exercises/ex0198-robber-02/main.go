// package main implements the *house robber* exercise from leetcode
// see: https://leetcode.com/problems/house-robber
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

func rob(nums []int) int {
	a, b := 0, 0
	for _, loot := range nums {
		a, b = b, max(b, a+loot)
		fmt.Printf("%v ", b)
	}
	return max(a, b)
}
