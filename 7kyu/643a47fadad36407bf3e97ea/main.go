package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(EncodeCd(5))   // PLLPPPPPP
	fmt.Println(EncodeCd(16))  // PPPPPLLLL
	fmt.Println(EncodeCd(63))  // PLPLPLPPP
	fmt.Println(EncodeCd(222)) // PPLPLPPLP
}

func EncodeCd(n uint8) string {
	var (
		str   = fmt.Sprintf("%08b", n)
		value = "P"
	)

	res := strings.Builder{}
	res.WriteString(value)

	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 49 {
			if value == "P" {
				value = "L"
			} else {
				value = "P"
			}
		}
		res.WriteString(value)
	}
	return res.String()
}

// https://www.codewars.com/kata/643a47fadad36407bf3e97ea
//
//
// "Pits" and "lands" are physical features on the surface of a CD that represent
// binary data. Pits are small depressions or grooves on the surface of the CD,
// while lands are flat areas between two adjacent pits.

// The pits and lands themselves do not directly represent the zeros and ones
// of binary data. Instead, Non-return-to-zero, inverted encoding is used: a change
// from pit to land or land to pit indicates a one, while no change indicates a zero.

// In this Kata, you should implement a function, that takes integer in range [0..255]
// (8 bits), and returns combination of pits and lands that encode the number.
// Result should have format of string: PLLPL... where P represents pit and L represents land.
// Combination should always start with pit. Numbers should be encoded in little-endian,
// which means you start from least significant bit.

// Example
// Input: 5

// Binary representation (8 bits): 00000101

// Output: PLLPPPPPP

// Explanation:

// Starts from P as per description
// Changes to L because first bit is 1
// Stays L because second bit is 0
// Changes to P because third bit is 1
// Stays P because fourth bit is 0
// Stays P till the end because all subsequent bits are 0
