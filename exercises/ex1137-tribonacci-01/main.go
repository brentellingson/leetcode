// package main solves the *nth- tribonacci number* exercise from leetcode
// see: https://leetcode.com/problems/n-th-tribonacci-number/
package main

import "fmt"

func main() {
	for i := range 38 {
		fmt.Printf("tribonacci(%v) = %v\n", i, tribonacci(i))
	}
}

// tribonacci implemented "dynamically"
func tribonacci(n int) int {
	t := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		t = append(t, t[i-3]+t[i-2]+t[i-1])
	}
	return t[n]
}
