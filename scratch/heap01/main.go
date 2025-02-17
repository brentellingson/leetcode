// implement a heap sort
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := makerandom(20, 5)
	fmt.Println(arr)
	heapify(arr)
	fmt.Println(arr)
	heapsort(arr)
	fmt.Println(arr)
}

func makerandom(count int, n int) []int {
	rslt := make([]int, 0, count)
	for range count {
		rslt = append(rslt, rand.Intn(n))
	}
	return rslt
}

func heapsort(heap []int) {
	for i := len(heap) - 1; i > 0; i-- {
		heap[0], heap[i] = heap[i], heap[0]
		siftdown(heap, 0, i)
	}
}

func heapify(arr []int) {
	// heapify all the subheaps, starting parent of the furthest leaf
	parent := (len(arr) - 2) / 2
	for root := parent; root >= 0; root-- {
		siftdown(arr, root, len(arr))
	}
}

func siftdown(arr []int, root int, end int) {
	for {
		left := root*2 + 1
		right := root*2 + 2
		if right < end && arr[right] > arr[root] && arr[right] > arr[left] {
			arr[right], arr[root] = arr[root], arr[right]
			root = right
			continue
		}
		if left < end && arr[left] > arr[root] {
			arr[left], arr[root] = arr[root], arr[left]
			root = left
			continue
		}
		break
	}
}
