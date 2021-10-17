package main

import (
	//"fmt"
	"math"
)

var dp [][]int

// func main() {
// 	wt := []int{1, 3, 4, 6}
// 	val := []int{1, 4, 5, 7}
// 	n := len(wt)
// 	weight := 7
// 	dp = make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		dp[i] = make([]int, weight+1)
// 	}
// 	for i := 0; i < n+1; i++ {
// 		for j := 0; j < weight+1; j++ {
// 			if i == 0 || j == 0 {
// 				dp[i][j] = 0
// 			} else {
// 				dp[i][j] = -1
// 			}
// 		}
// 	}
// 	fmt.Println("max profit ", knapsack_tw(wt, val, weight, n))
// }

func knapsack_dp(wt []int, val []int, weight int, n int) int {
	if weight == 0 || n == 0 {
		return 0
	}
	if dp[n][weight] != -1 {
		return dp[n][weight]
	}
	if wt[n-1] <= weight {
		dp[n][weight] = int(math.Max(float64(val[n-1]+knapsack_dp(wt, val, weight-wt[n-1], n-1)), float64(knapsack_dp(wt, val, weight, n-1))))
	} else {
		dp[n][weight] = knapsack_dp(wt, val, weight, n-1)
	}
	return dp[n][weight]
}

func knapsack_tw(wt []int, val []int, weight int, n int) int {
	for i := 1; i <= n; i++ {
		for j := 1; j <= weight; j++ {
			if wt[i-1] <= j {
				dp[i][j] = int(math.Max(float64(val[i-1]+dp[i-1][j-wt[i-1]]), float64(dp[i-1][j])))
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][weight]
}
