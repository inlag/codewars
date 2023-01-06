package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(SumMix([]any{9, 3, "7", "3"})) //22
}

func SumMix(arr []any) int {
	var response int
	for _, val := range arr {
		var number int
		valStr, okStr := val.(string)
		if okStr {
			number, _ = strconv.Atoi(valStr)
		} else {
			number, _ = val.(int)
		}
		response += number
	}
	return response
}

//Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.
//
//Return your answer as a number.
