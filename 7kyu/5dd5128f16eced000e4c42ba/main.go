package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solve("a"))        // 97
	fmt.Println(Solve("aa"))       // 97
	fmt.Println(Solve("bcd"))      // 98
	fmt.Println(Solve("axyzxyz"))  // 120
	fmt.Println(Solve("dcbadcba")) // 97
	fmt.Println(Solve("aabccc"))   // 99
}

func Solve(s string) rune {
	if len(s) == 1 {
		return rune(s[0])
	}

	var (
		maxIndex      = 0
		minRune  rune = 0
	)
	for index, val := range s {
		for i := len(s) - 1; i >= index; i-- {
			if val == rune(s[i]) {
				diff := i - index
				// fmt.Println("word : ", string(s[i]), " diff: ", diff)
				if diff > maxIndex {
					maxIndex = diff
					minRune = val
				} else if diff == maxIndex && minRune > val {
					maxIndex = diff
					minRune = val
				}
			}
		}
	}

	if minRune == 0 {
		return rune(s[0])
	}
	return minRune
}

// https://www.codewars.com/kata/5dd5128f16eced000e4c42ba
//
//
// In this Kata, you will be given a string and your task is to return
// the most valuable character. The value of a character is the difference
// between the index of its last occurrence and the index of its first occurrence.
// Return the character that has the highest value. If there is a tie, return
// the alphabetically lowest character. [For Golang return rune]

// All inputs will be lower case.

// For example:
// solve('a') = 'a'
// solve('ab') = 'a'. Last occurrence is equal to first occurrence of each character.
// 	Return lexicographically lowest.
// solve("axyzxyz") = 'x'
// More examples in test cases. Good luck!
