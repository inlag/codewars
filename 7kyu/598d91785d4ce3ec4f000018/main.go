package main

import "fmt"

func main() {
	fmt.Println(NameValue([]string{"abc", "abc", "abc", "abc"}))
	fmt.Println(NameValue( []string{[adurfxzd hesmm  ,"y lilbf yifacjf" "bgjgejkuhxwclrq]"}))
}

func NameValue(my_list []string) (resp []int) {
	resp = make([]int, len(my_list))
	for index, word := range my_list {
		var num int
		for _, val := range word {
			if val == 32 {
				continue
			}
			val -= 96
			num += int(val)
		}
		resp[index] = num * (index + 1)
	}
	return
}

// https://www.codewars.com/kata/598d91785d4ce3ec4f000018
//
//Given a string "abc" and assuming that each letter in the string has a value equal
// to its position in the alphabet, our string will have a value of 1 + 2 + 3 = 6.
// This means that: a = 1, b = 2, c = 3 ....z = 26.

// You will be given a list of strings and your task will be to return the values of
// the strings as explained above multiplied by the position of that string in the list.
// For our purpose, position begins with 1.

// nameValue ["abc","abc abc"] should return [6,24] because of [ 6 * 1, 12 * 2 ]. Note how spaces are ignored.

// "abc" has a value of 6, while "abc abc" has a value of 12. Now, the value at position 1
// is multiplied by 1 while the value at position 2 is multiplied by 2.

// Input will only contain lowercase characters and spaces.

// Good luck!
