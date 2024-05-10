package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func Capitalize(st string, arr []int) string {
	res := []byte(st)
	stCount := len(st)

	for i := 0; i < len(arr); i++ {
		if arr[i] > stCount {
			continue
		}

		tt := strings.ToUpper(string(st[arr[i]]))
		res[arr[i]] = *unsafe.StringData(tt)
	}
	return string(res)
}

func main() {
	fmt.Println(Capitalize("abcdef", []int{1, 2, 5, 100}))
}

// https://www.codewars.com/kata/59cfc09a86a6fdf6df0000f1
//
//Given a string and an array of integers representing indices, capitalize all letters at the given indices.

// For example:

// capitalize("abcdef",[1,2,5]) = "aBCdeF"
// capitalize("abcdef",[1,2,5,100]) = "aBCdeF". There is no index 100.
// The input will be a lowercase string with no spaces and an array of digits.

// Good luck!
