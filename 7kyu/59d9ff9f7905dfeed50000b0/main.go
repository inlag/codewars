package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solve([]string{"abode", "ABc", "xyzD"}))
}

var alphabet = []string{
	"a", "b", "c", "d", "e", "f",
	"g", "h", "i", "j", "k", "l",
	"m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x",
	"y", "z",
}

func solve(slice []string) []int {
	var res = make([]int, len(slice))
	for iWord, word := range slice {
		for iLetter, letter := range strings.ToLower(word) {
			if string(letter) == alphabet[iLetter] {
				res[iWord] = res[iWord] + 1
			}
		}
	}
	return res
}
