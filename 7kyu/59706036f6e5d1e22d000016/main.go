package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(WordsToMarks("attitude"))
	fmt.Println(WordsToMarks("friends"))
}

func WordsToMarks(s string) int {
	var res int32
	for _, val := range strings.ToLower(s) {
		res += val - 96
	}
	return int(res)
}

//If　a = 1, b = 2, c = 3 ... z = 26
//
//Then l + o + v + e = 54
//
//and f + r + i + e + n + d + s + h + i + p = 108
//
//So friendship is twice as strong as love :-)
//
//Your task is to write a function which calculates the value of a word based
//off the sum of the alphabet positions of its characters.
//
//The input will always be made of only lowercase letters and will never be
//empty.
