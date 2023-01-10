package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps")) // => 3
	fmt.Println(FindShort("Let's travel abroad shall we"))                        // => 2
}

func FindShort(s string) int {
	arr := strings.Split(s, " ")

	var resInt int = utf8.RuneCountInString(arr[0])
	for _, word := range arr[0:] {
		count := utf8.RuneCountInString(word)
		if resInt > count {
			resInt = count
		}
	}
	return resInt
}

//Simple, given a string of words, return the length of the shortest word(s).
//
//String will never be empty and you do not need to account for different data types.
