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

// fiboonacci implemented recursively with a higher-order memoize function
// Use a variable declaration so it can be bound to to the memoized version
// Define it in an init() block because go won't let use recurse definitions outsize a block
var fibonacci func(n int) int

func init() {
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
	fibonacci = memoize(fibonacci)
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
