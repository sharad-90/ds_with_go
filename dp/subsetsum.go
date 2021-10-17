package main

import "fmt"

var dp_sum [][]bool

func main() {
	arr := []int{1, 5, 11, 5}
	sum := 11
	fmt.Println(subSetSum(arr, sum, len(arr)))
	dp_sum = make([][]bool, len(arr)+1)
	for i := 0; i < len(arr)+1; i++ {
		dp_sum[i] = make([]bool, sum+1)
	}
	for i := 0; i < len(arr)+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 {
				dp_sum[i][j] = false
			}
			if j == 0 {
				dp_sum[i][j] = true
			}
		}
	}
	//fmt.Println(equalSum(arr))
}

func subSetSum(arr []int, sum, n int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 {
		return false
	}
	if arr[n-1] <= sum {
		return subSetSum(arr, sum-arr[n-1], n-1) || subSetSum(arr, sum, n-1)
	} else {
		return subSetSum(arr, sum, n-1)
	}
}

func subSetSum_dp(arr []int, sum, n int) bool {
	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[n-1] <= sum {
				dp_sum[i][j] = dp_sum[i-1][sum-arr[n-1]] || dp_sum[i-1][sum]
			} else {
				dp_sum[i][j] = dp_sum[i-1][sum]
			}
		}
	}
	return dp_sum[n][sum]
}
/*
func equalSum(arr []int) bool {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
	}
	if n%2 != 0 {
		return false
	} else {
		return subSetSum_dp(arr, sum/2, n)
	}
}
*/
