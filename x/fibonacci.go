package x

// fiboonacci implemented recursively with a higher-order memoize function
// Use a variable declaration so it can be bound to to the memoized version
// Define it in an init() block because go won't let use recurse definitions outsize a block
var Fibonacci func(n int) int

func init() {
	Fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
	Fibonacci = Memoize(Fibonacci)
}
