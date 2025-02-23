// package main implements the *climbing stairs* exercise from leetcode
// see: https://leetcode.com/problems/climbing-stairs/
package main

import "fmt"

func main() {
	for i := range 45 {
		fmt.Printf("climb starts %v = %v\n", i+1, climbStairs(i+1))
	}
}

// climbStairs(i) is fibonacci(i+1)  (see sanskrit prosody)
func climbStairs(n int) int {
	return fibonacci(n + 1)
}

// fibonacci implemented recursively. This approach is very slow due to repeated calculations.
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
