package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SortVowels("Codewars"))
	fmt.Println(SortVowels("Rnd Te5T"))
}

func SortVowels(s string) string {
	if s == "" {
		return ""
	}
	const words = "a,e,i,o,u"
	response := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if strings.Contains(words, strings.ToLower(string(s[i]))) {
			response.WriteString("|" + string(s[i]) + "\n")
			continue
		}
		response.WriteString(string(s[i]) + "|" + "\n")
	}
	return response.String()[:response.Len()-1]
}

// https://www.codewars.com/kata/59e49b2afc3c494d5d00002a
//
// Sort the Vowels!
// In this kata, we want to sort the vowels in a special format.

// Task
// Write a function which takes a input string s and return a string in the following way:

//                   C|                          R|
//                   |O                          n|
//                   D|                          d|
//    "CODEWARS" =>  |E       "Rnd Te5T"  =>      |
//                   W|                          T|
//                   |A                          |e
//                   R|                          5|
//                   S|                          T|

// Notes:

// List all the Vowels on the right side of |
// List every character except Vowels on the left side of |
// for the purpose of this kata, the vowels are : a e i o u
// Return every character in its original case
// Each line is seperated with \n
// Invalid input ( undefined / null / integer ) should return an empty string
