// package main implements heap sort
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sort(10)
	sort(20)
	sort(30)
	sort(40)
	sort(60)
}

func sort(count int) {
	arr := []int{}
	cmp := func(a, b int) int {
		return a - b
	}
	for range count {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	heapsort(arr, cmp)
	fmt.Printf("%v, swaps=%.2f, comps=%.2f\n", arr, float64(swaps)/float64(count), float64(comps)/float64(count))

	for i := range arr {
		arr[i] = i
	}
	heapsort(arr, cmp)
	fmt.Printf("%v, swaps=%.2f, comps=%.2f\n", arr, float64(swaps)/float64(count), float64(comps)/float64(count))

	for i := range arr {
		arr[len(arr)-i-1] = i
	}
	heapsort(arr, cmp)
	fmt.Printf("%v, swaps=%.2f, comps=%.2f\n", arr, float64(swaps)/float64(count), float64(comps)/float64(count))
}

func makerandom(count int, n int) []int {
	rslt := make([]int, 0, count)
	for range count {
		rslt = append(rslt, rand.Intn(n))
	}
	return rslt
}

var swaps int
var comps int

func heapsort[S ~[]E, E any](arr S, cmp func(a, b E) int) {
	swaps, comps = 0, 0
	swap := func(a, b int) {
		swaps++
		arr[a], arr[b] = arr[b], arr[a]
	}
	less := func(a, b int) bool {
		comps++
		return cmp(arr[a], arr[b]) < 0
	}

	start, end := len(arr)/2, len(arr)
	for end > 1 {
		if start > 0 {
			start--
		} else {
			end--
			swap(0, end)
		}
		root := start
		for root*2+1 < end {
			left, right := root*2+1, root*2+2
			if right < end && less(root, right) && less(left, right) {
				swap(root, right)
				root = right
				continue
			}
			if left < end && less(root, left) {
				swap(root, left)
				root = left
				continue
			}
			break
		}
	}
}
