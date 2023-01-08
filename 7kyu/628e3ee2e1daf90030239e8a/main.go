package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	fmt.Println(Interlockable(9, 4)) // => true
	fmt.Println(Interlockable(3, 6)) // => false
}

func Interlockable(a, b uint64) bool {
	strA := reverse(strconv.FormatUint(a, 2))
	strB := reverse(strconv.FormatUint(b, 2))

	var arr = strB
	var arr2 = strA
	if utf8.RuneCountInString(strB) > utf8.RuneCountInString(strA) {
		arr = strA
		arr2 = strB
	}
	for i, v := range arr {
		if string(v) == "1" {
			if string(arr2[i]) == "1" {
				return false
			}
		}

	}

	return true
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

//Write a function that checks if two non-negative integers make an "interlocking binary pair".
//
//Interlock ?
//numbers can be interlocked if their binary representations have no 1's in the same place
//comparisons are made by bit position, starting from right to left (see the examples below)
//when representations are of different lengths, the unmatched left-most bits are ignored
//Examples
//a = 9, b = 4
//Stacking representations shows how they can interlock.
//
//9    1001
//4     100
//Here, no 1's share any position, so the function returns true.
//
//a = 3, b = 6
//These representations do not interlock.
//
//3      11
//6     110
//Finding they both have a 1 in the same position, the function returns false.
//
//Input
//Two non-negative integers.
//
//Output
//Boolean true or false whether or not these integers are interlockable.
