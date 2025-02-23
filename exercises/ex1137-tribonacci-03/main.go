// package main solves the *nth- tribonacci number* exercise from leetcode
// see: https://leetcode.com/problems/n-th-tribonacci-number/
package main

import "fmt"

func main() {
	for i := range 38 {
		fmt.Printf("tribonacci(%v) = %v\n", i, tribonacci(i))
	}
}

// tribonacci implemented with higher-order memoization
func tribonacci(n int) int {
	return tribo(n)
}

var tribo func(n int) int

func init() {
	tribo = func(n int) int {
		switch n {
		case 0:
			return 0
		case 1, 2:
			return 1
		default:
			return tribo(n-3) + tribo(n-2) + tribo(n-1)
		}
	}
	tribo = memoize(tribo)
}

func memoize[U comparable, V any](f func(U) V) func(U) V {
	cache := make(map[U]V)
	return func(x U) V {
		if _, ok := cache[x]; !ok {
			cache[x] = f(x)
		}
		return cache[x]
	}
}
