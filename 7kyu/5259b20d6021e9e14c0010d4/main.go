package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog."))
}

func ReverseWords(str string) string {
	var (
		words = strings.Split(str, " ")
		res   = make([]string, len(words))
	)
	for index, word := range words {
		res[index] = reversWord(word)
	}

	return strings.Join(res, " ")
}

func reversWord(w string) string {
	var response = make([]byte, len(w))
	for i, ii := 0, len(w)-1; ii >= 0; i, ii = i+1, ii-1 {
		response[ii] = w[i]
	}
	return string(response)
}

// https://www.codewars.com/kata/5259b20d6021e9e14c0010d4
//Complete the function that accepts a string parameter,
// and reverses each word in the string. All spaces in the string should be retained.

// Examples
// "This is an example!" ==> "sihT si na !elpmaxe"
// "double  spaces"      ==> "elbuod  secaps"
