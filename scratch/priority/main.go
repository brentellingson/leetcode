package main

import (
	"fmt"
	"math/rand"
)

func main() {
	compare := func(a, b int) int {
		return a - b
	}
	pq := &PriorityQueue[int]{Cmp: compare}
	for range 10 {
		pq.Push(rand.Intn(100))
	}
	for range 10 {
		pq.Push(rand.Intn(100))
		fmt.Println(pq.Pull())
	}
	for !pq.IsEmpty() {
		fmt.Println(pq.Pull())
	}
}

func makerandom(count int, n int) []int {
	rslt := make([]int, 0, count)

	for range count {
		rslt = append(rslt, rand.Intn(n))
	}
	return rslt
}

// priority queue
type PriorityQueue[T any] struct {
	Heap []T
	Cmp  func(a, b T) int
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.Heap) == 0
}

func (pq *PriorityQueue[T]) Push(elem T) {
	pq.Heap = append(pq.Heap, elem)
	pq.shiftup()
}

func (pq *PriorityQueue[T]) Pull() (T, bool) {
	var zero T

	if len(pq.Heap) == 0 {
		return zero, false
	}
	rslt := pq.Heap[0]
	tail := pq.Heap[len(pq.Heap)-1]
	pq.Heap = pq.Heap[:len(pq.Heap)-1]
	if len(pq.Heap) > 0 {
		pq.Heap[0] = tail
		pq.shiftdown()
	}
	return rslt, true
}

// shiftup fixes a heap where the last element may not be the smallest.
// if pulls the last element up the heap to its proper position
func (pq *PriorityQueue[T]) shiftup() {
	parent := func(child int) int {
		return (child - 1) / 2
	}
	compare := func(a, b int) int {
		return pq.Cmp(pq.Heap[a], pq.Heap[b])
	}
	swap := func(a, b int) {
		pq.Heap[a], pq.Heap[b] = pq.Heap[b], pq.Heap[a]
	}

	child := len(pq.Heap) - 1
	for child > 0 && compare(parent(child), child) < 0 {
		swap(parent(child), child)
		child = parent(child)
	}
}

// shiftdown fixes a heap where the first element may not be the largest.
// it pushes the first element down the heap to its proper position
func (pq *PriorityQueue[T]) shiftdown() {
	compare := func(a, b int) int {
		return pq.Cmp(pq.Heap[a], pq.Heap[b])
	}
	swap := func(a, b int) {
		pq.Heap[a], pq.Heap[b] = pq.Heap[b], pq.Heap[a]
	}

	root := 0
	for {
		left := root*2 + 1
		right := root*2 + 2
		if right < len(pq.Heap) && compare(root, right) < 0 && compare(left, right) < 0 {
			swap(root, right)
			root = right
		} else if left < len(pq.Heap) && compare(root, left) < 0 {
			swap(root, left)
			root = left
		} else {
			break
		}
	}
}
