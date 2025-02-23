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

// fibonacci without recursion (sort of)
func fibonacci(n int) int {
	a, b := 0, 1
	for range n {
		a, b = b, a+b
	}
	return a
}
