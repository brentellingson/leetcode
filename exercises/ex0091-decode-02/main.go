// package main implements the *decode ways* exercise from leetcode
// see: https://leetcode.com/problems/decode-ways
package main

import "fmt"

func main() {
	print("0")
	print("1")
	print("2")
	print("3")
	print("12")
	print("226")
	print("06")
	print("911")
	print("26")
	print("27")
	print("262626")
}

func print(s string) {
	fmt.Printf("numDeconds(%v) = %v\n", s, numDecodings(s))
}

func numDecodings(s string) int {
	var num func(s string) int
	num = func(s string) int {
		if len(s) == 0 {
			return 1
		}
		if s[0] == '0' {
			return 0
		}
		if len(s) > 1 && s[0] == '1' {
			return num(s[1:]) + num(s[2:])
		}
		if len(s) > 1 && s[0] == '2' && s[1] <= '6' {
			return num(s[1:]) + num(s[2:])
		}
		return num(s[1:])
	}
	num = memoize(num)

	return num(s)
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
