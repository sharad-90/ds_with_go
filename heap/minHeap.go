package main

import (
	"fmt"
	"math"
)

type minHeap struct {
	size     int
	capacity int
	arr      []int
}

func newMinHeap(c int) *minHeap {
	heap := new(minHeap)
	heap.size = 0
	heap.capacity = c
	heap.arr = make([]int, c, c)
	return heap
}

func left(i int) int {
	return (2*i + 1)
}

func right(i int) int {
	return (2*i + 2)
}

func parent(i int) int {
	return (i - 1) / 2
}

func (heap *minHeap) insert(i int) {
	if heap.size == heap.capacity {
		return
	}
	heap.size = heap.size + 1
	heap.arr[heap.size-1] = i
	for i := (heap.size - 1); i != 0 && heap.arr[parent(i)] > heap.arr[i]; {
		heap.swap(i, parent(i))
		i = parent(i)
	}
}

func (heap *minHeap) minHeapify(i int) {
	lt := left(i)
	rt := right(i)
	smallest := i
	if heap.size > lt && heap.arr[lt] < heap.arr[i] {
		smallest = lt
	}
	if heap.size > rt && heap.arr[rt] < heap.arr[smallest] {
		smallest = rt
	}
	if smallest != i {
		heap.swap(i, smallest)
		heap.minHeapify(smallest)
	}
}

func (heap *minHeap) extractMin() int {
	if heap.size == 0 {
		return math.MinInt8
	}
	if heap.size == 1 {
		heap.size = heap.size - 1
		return heap.arr[heap.size]
	}
	heap.swap(0, heap.size-1)
	heap.size = heap.size - 1
	heap.minHeapify(0)
	return heap.arr[heap.size]
}

func (heap *minHeap) decreaseKey(i, x int) {
	heap.arr[i] = x
	for i != 0 && heap.arr[i] < heap.arr[parent(i)] {
		heap.swap(i, parent(i))
		i = parent(i)
	}
}

func (heap *minHeap) deleteKey(i int) {
	heap.decreaseKey(i, math.MinInt8)
	heap.extractMin()
}

func (heap *minHeap) buildMinHeap() {
	for i := (heap.size - 2) / 2; i >= 0; i-- {
		heap.minHeapify(i)
	}
}

func (heap *minHeap) swap(i, j int) {
	tmp := heap.arr[i]
	heap.arr[i] = heap.arr[j]
	heap.arr[j] = tmp
}

func maxHeap(arr []int) {
	n := len(arr)
	for i := (n - 2) / 2; i >= 0; i-- {
		maxHeapify(arr, n, i)
	}
}

func maxHeapify(arr []int, n, i int) {
	lt := 2*i + 1
	rt := 2*i + 2
	largest := i
	if lt < n && arr[largest] < arr[lt] {
		largest = lt
	}
	if rt < n && arr[largest] < arr[rt] {
		largest = rt
	}
	if largest != i {
		tmp := arr[i]
		arr[i] = arr[largest]
		arr[largest] = tmp
		maxHeapify(arr, n, largest)
	}
}

func heapSort(arr []int) {
	maxHeap(arr)
	n := len(arr)
	for i := n - 1; i >= 1; i-- {
		tmp := arr[i]
		arr[i] = arr[0]
		arr[0] = tmp
		maxHeapify(arr, i, 0)
	}
}

func printKLargest(arr []int, k int) {
	karr := make([]int, k)
	for i := 0; i < k; i++ {
		karr[i] = arr[i]
	}
	fmt.Println("k elements are ", karr[0])
	buildMinHeap(karr)

	for j := k + 1; j < len(arr); j++ {
		if arr[j] > karr[0] {
			karr[0] = arr[j]
			buildMinHeap(karr)
		}
	}
	fmt.Println("max k elements are ", karr)
}
