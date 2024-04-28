package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solve("codewars", 1, 5))      // cawedors
	fmt.Println(Solve("codingIsFun", 2, 100)) // cawedors
}

func Solve(s string, a, b int) string {
	if b >= len(s) {
		bytes := []byte(s[a:])
		str := reverse(bytes)
		return s[:a] + str
	}
	bytes := []byte(s[a : b+1])
	str := reverse(bytes)
	return s[:a] + str + s[b+1:]
}

func reverse(s []byte) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

// https://www.codewars.com/kata/5a8d1c82373c2e099d0000ac
//
// In this Kata, you will be given a string and two indexes (a and b).
// Your task is to reverse the portion of that string between those two indices inclusive.

// solve("codewars",1,5) = "cawedors" -- elements at index 1 to 5 inclusive
// are "odewa". So we reverse them.
// solve("cODEWArs", 1,5) = "cAWEDOrs" -- to help visualize.
// Input will be lowercase and uppercase letters only.

// The first index a will always be lower that than the string length;
// the second index b can be greater than the string length.
// More examples in the test cases. Good luck!
