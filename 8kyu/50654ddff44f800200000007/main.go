package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(Solution("13", "200")) //1320013
	fmt.Println(Solution("45", "1"))   //1451
}

func Solution(a, b string) string {
	var min, max = a, b
	if utf8.RuneCount([]byte(a)) > utf8.RuneCount([]byte(b)) {
		min, max = b, a
	}

	return min + max + min
}

//Given 2 strings, a and b, return a string of the form short+long+short, with
//the shorter string on the outside and the longer string on the inside. The
//strings will not be the same length, but they may be empty ( zero length ).
//
//For example: (Input1, Input2) --> output
//
//("1", "22") --> "1221"
//("22", "1") --> "1221"
