package main

import "fmt"

func main() {
	fmt.Println(StantonMeasure([]int{1, 4, 3, 2, 1, 2, 3, 2}))
	fmt.Println(StantonMeasure([]int{1, 4, 3, 0, 1, 9, 3, 6}))
}

func StantonMeasure(arr []int) int {
	var tmp = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		tmp[arr[i]]++
	}
	return tmp[tmp[1]]
}

// https://www.codewars.com/kata/59a1cdde9f922b83ee00003b
//
// The Stanton measure of an array is computed as follows: count the number of occurences
// for value 1 in the array. Let this count be n. The Stanton measure is the number of times
// that n appears in the array.

// Write a function which takes an integer array and returns its Stanton measure.

// Examples
// The Stanton measure of [1, 4, 3, 2, 1, 2, 3, 2] is 3, because 1 occurs 2 times in the array and 2 occurs 3 times.

// The Stanton measure of [1, 4, 1, 2, 11, 2, 3, 1] is 1, because 1 occurs 3 times in the array and 3 occurs 1 time.
