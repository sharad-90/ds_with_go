package main

import "fmt"

func main() {
	heap := newMinHeap(5)
	heap.insert(100)
	heap.insert(10)
	heap.insert(5)
	heap.insert(101)
	heap.insert(2)
	fmt.Println("Hello world", heap.arr[0:heap.size])
	heap.deleteKey(2)
	fmt.Println("Hello world", heap.arr[0:heap.size])
	arr_1 := []int{10, 0, 5, 20, 3}
	heapSort(arr_1)
	fmt.Println(arr_1)
	arr_2 := []int{10, 2, 100, 1, 4}
	buildMinHeap(arr_2)
	fmt.Println(arr_2)
	arr_3 := []int{10, 2, 100, 1, 4}
	printKLargest(arr_3, 3)
}
