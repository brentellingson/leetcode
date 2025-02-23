package x

func Memoize[K comparable, V any](f func(k K) V) func(k K) V {
	memo := make(map[K]V)
	return func(k K) V {
		if _, ok := memo[k]; !ok {
			memo[k] = f(k)
		}
		return memo[k]
	}
}

func Memoize2[K1 comparable, K2 comparable, V any](f func(k1 K1, k2 K2) V) func(k1 K1, k2 K2) V {
	type key2 struct {
		k1 K1
		k2 K2
	}
	memo := make(map[key2]V)
	return func(k1 K1, k2 K2) V {
		if _, ok := memo[key2{k1, k2}]; !ok {
			memo[key2{k1, k2}] = f(k1, k2)
		}
		return memo[key2{k1, k2}]
	}
}
