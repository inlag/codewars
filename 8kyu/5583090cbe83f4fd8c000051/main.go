package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println(Digitize(0))
	fmt.Println(Digitize(35231))
}

func Digitize(n int) []int {
	str := strconv.Itoa(n)

	lengthStr := utf8.RuneCountInString(str)
	var response = make([]int, lengthStr, lengthStr)
	for i, val := range str {
		valInt, _ := strconv.Atoi(string(val))
		response[(lengthStr-1)-i] = valInt
	}
	return response
}

//Given a random non-negative number, you have to return the digits of this number within an array in reverse order.
