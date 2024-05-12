package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SumOfIntegersInString("V6XPJSZYH.V590KBI.TW4KZ5ER1PDMBAHMCGQRMMM52G-CSYGALCSZ5QEY1FLNMKYW0TCP3EHTRE9ZXFU58624AS6E0YC.CTDX45"))
	fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))
	fmt.Println(SumOfIntegersInString("Our company made approximately 1 million in gross revenue last quarter."))
	fmt.Println(SumOfIntegersInString("12.4"))
}

func SumOfIntegersInString(strng string) (resp int) {
	const numbers = "0123456789"
	for i := 0; i < len(strng); i++ {
		if strings.Contains(numbers, string(strng[i])) {
			var numberStr string
			for ii := i; ii < len(strng); ii++ {
				i = ii
				if strings.Contains(numbers, string(strng[ii])) {
					numberStr += string(strng[ii])
				} else {
					break
				}
			}
			number, _ := strconv.Atoi(string(numberStr))
			resp += number
		}
	}
	return
}

// https://www.codewars.com/kata/598f76a44f613e0e0b000026
//
// Your task in this kata is to implement a function that calculates the sum of the integers inside a string.
// For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog",
// the sum of the integers is 3635.

// Note: only positive integers will be tested.
