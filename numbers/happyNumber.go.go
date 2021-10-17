package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(isMagicNumber(7))
}

func isMagicNumber(i int) bool {
	result := squareSum(i)
	if result == 1 {
		return true
	} else {
		return false
	}
}

func squareSum(num int) int {
	strNum := strconv.Itoa(num)
	var result int
	if num == 1 || num == 4 {
		return num
	} else {
		for i := 0; i < len(strNum); i++ {
			digit, _ := strconv.Atoi(string(strNum[i]))
			result = result + digit*digit
		}
	}
	return squareSum(result )
}
