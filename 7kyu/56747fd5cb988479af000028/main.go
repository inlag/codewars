package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(GetMiddle("hello"))
	fmt.Println(GetMiddle("test"))
	fmt.Println(GetMiddle("testing"))
	fmt.Println(GetMiddle("middle"))
	fmt.Println(GetMiddle("A"))
}
func GetMiddle(s string) string {
	countRunes := utf8.RuneCountInString(s)
	if countRunes%2 == 0 {
		return s[countRunes/2-1 : countRunes/2+1]
	}
	return string(s[countRunes/2])
}

// https://www.codewars.com/kata/56747fd5cb988479af000028
//
// You are going to be given a word. Your job is to return the middle character of the word.
// If the word's length is odd, return the middle character.
//  If the word's length is even, return the middle 2 characters.

// #Examples:

// Kata.getMiddle("test") should return "es"

// Kata.getMiddle("testing") should return "t"

// Kata.getMiddle("middle") should return "dd"

// Kata.getMiddle("A") should return "A"
