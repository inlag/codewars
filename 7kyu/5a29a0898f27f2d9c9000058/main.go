package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(Solve("Codewars@codewars123.com"))
}

func Solve(s string) []int {
	var (
		upper, lower, numbers, special int
	)

	for _, letter := range s {
		if unicode.IsUpper(letter) {
			upper++
		} else if unicode.IsLower(letter) {
			lower++
		} else if unicode.IsNumber(letter) {
			numbers++
		} else {
			special++
		}
	}

	return []int{upper, lower, numbers, special}
}

//In this Kata, you will be given a string and your task will be to return a
//list of ints detailing the count of uppercase letters, lowercase, numbers and
//special characters, as follows.
//
//Solve("*'&ABCDabcde12345") = [4,5,5,3].
//--the order is: uppercase letters, lowercase, numbers and special characters.
//More examples in the test cases.
//
//Good luck!
