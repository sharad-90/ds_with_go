package main

import (
	//"fmt"
	"math"
)

// func main() {
// 	wt := []int{1, 3, 4, 6}
// 	val := []int{1, 4, 5, 7}
// 	weight := 7
// 	n := len(wt)
// 	fmt.Println("max profit ", knapsack(wt, val, weight, n))
// }

func knapsack(wt []int, val []int, weight int, n int) int {
	if weight == 0 || n == 0 {
		return 0
	}
	if wt[n-1] <= weight {
		return int(math.Max(float64(val[n-1]+knapsack(wt, val, weight-wt[n-1], n-1)), float64(knapsack(wt, val, weight, n-1))))
	} else {
		return knapsack(wt, val, weight, n-1)
	}

}
