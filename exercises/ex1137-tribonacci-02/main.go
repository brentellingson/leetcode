// package main solves the *nth- tribonacci number* exercise from leetcode
// see: https://leetcode.com/problems/n-th-tribonacci-number/
package main

import "fmt"

func main() {
	for i := range 38 {
		fmt.Printf("tribonacci(%v) = %v\n", i, tribonacci(i))
	}
}

// tribonacci implemented iteratively
func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	for range n {
		a, b, c = b, c, a+b+c
	}
	return a
}
