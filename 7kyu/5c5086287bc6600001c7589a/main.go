package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(IsNegativeZero(math.Copysign(0, -1)))
	fmt.Println(IsNegativeZero(0.0))
}

func IsNegativeZero(n float64) bool {
	return strings.EqualFold("-0", fmt.Sprintf("%v", n))
}

// https://www.codewars.com/kata/5c5086287bc6600001c7589a
//
// There exist two zeroes: +0 (or just 0) and -0.

// Write a function that returns true if the input number is -0 and false otherwise (True and False for Python).
