package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(ReverseLetters("krishan"))
}

func ReverseLetters(s string) string {
	var res strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if unicode.IsLetter(r) {
			res.WriteRune(r)
		}
	}
	return res.String()
}

// https://www.codewars.com/kata/58b8c94b7df3f116eb00005b
//
// Given a string str, reverse it and omit all non-alphabetic characters.
//
// Example
// For str = "krishan", the output should be "nahsirk".
//
// For str = "ultr53o?n", the output should be "nortlu".
//
// Input/Output
// [input] string str
// A string consists of lowercase latin letters, digits and symbols.
//
// [output] a string
