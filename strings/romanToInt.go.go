package main

import "fmt"

func main(){
fmt.Println(romanToInt("IV"))
}

func romanToInt(s string) int {
    numMap := make(map[string]int)
	numMap["I"] = 1
	numMap["V"] = 5
	numMap["X"] = 10
	numMap["L"] = 50
	numMap["C"] = 100
	numMap["D"] = 500
	numMap["M"] = 1000
    var resultant int
	if len(s) >= 1 && len(s) <= 15 {
		for i := 0; i < len(s); i++ {
			if i+1 < len(s) && numMap[string(s[i+1])] > numMap[string(s[i])] {
				resultant = resultant + numMap[string(s[i+1])] - numMap[string(s[i])]
				i = i + 1
			} else {
				resultant = resultant + numMap[string(s[i])]
			}
		}
	}
    return resultant
}