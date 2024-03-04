package main

import "fmt"

func main() {
	fmt.Println(Smaller([]int{5, 4, 3, 2, 1}))
}

func Smaller(arr []int) []int {
	var result = make([]int, len(arr))

	for index, value := range arr {
		var count int
		for _, valueNew := range arr[index:] {
			if value > valueNew {
				count++
			}
		}
		result[index] = count
	}
	return result
}

// https://www.codewars.com/kata/56a1c074f87bc2201200002e
//
// Write a function that given, an array arr, returns an array containing at
// each index i the amount of numbers that are smaller than arr[i] to the right.
//
// For example:
//
// * Input [5, 4, 3, 2, 1] => Output [4, 3, 2, 1, 0]
// * Input [1, 2, 0] => Output [1, 1, 0]
