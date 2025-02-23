// package main solves the *perfect squares* exercise from leetcode
// see: https://leetcode.com/problems/perfect-squares/
package main

import (
	"fmt"
	"slices"
)

func main() {
	for i := range 1000 {
		fmt.Printf("squares(%v) = %v = %v\n", i+1, squares(i+1), sum(squares(i+1)))
	}
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

// numsquares implemented with higher-order memoization
func numSquares(n int) int {
	return len(squares(n))
}

var squares func(n int) []int

func init() {
	squares = func(n int) []int {
		if n == 0 {
			return nil
		}
		rslt := append(squares(n-1), 1)
		for i := 2; i*i <= n; i++ {
			steps := append(squares(n-i*i), i*i)
			if len(steps) < len(rslt) {
				rslt = steps
			}
		}
		// clone so we don't accidentally alias a result
		return slices.Clone(rslt)
	}
	squares = memoize(squares)
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
