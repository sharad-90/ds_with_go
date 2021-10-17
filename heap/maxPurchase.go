package main

import (
	"fmt"
	"math"
)

//{10, 2, 100, 1, 4}
func buildMinHeap(arr []int) {
	n := len(arr)
	sum, res := 18, 0
	for i := (n - 2) / 2; i >= 0; i-- {
		minHeapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		min := extract(arr, i+1)
		if sum >= min {
			sum = sum - min
			res++
		}
	}
	fmt.Println("Result is", res)
}

func minHeapify(arr []int, n, i int) {
	lt := 2*i + 1
	rt := 2*i + 2
	smallest := i
	if lt < n && arr[lt] < arr[i] {
		smallest = lt
	}
	if rt < n && arr[rt] < arr[smallest] {
		smallest = rt
	}
	if smallest != i {
		tmp := arr[i]
		arr[i] = arr[smallest]
		arr[smallest] = tmp
		minHeapify(arr, n, smallest)
	}
}

func extract(arr []int, size int) int {
	if size == 0 {
		return math.MinInt8
	}
	if size == 1 {
		return arr[0]
	}
	tmp := arr[0]
	arr[0] = arr[size-1]
	arr[size-1] = tmp
	minHeapify(arr, size-1, 0)
	return arr[size-1]
}
