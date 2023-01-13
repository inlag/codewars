package main

import "fmt"

func main() {
	fmt.Println(HasUniqueChar("  nAa"))
	fmt.Println(HasUniqueChar("abcde"))
}

func HasUniqueChar(str string) bool {
	var arr = make(map[int32]bool, 0)
	for _, letter := range str {
		if ok, _ := arr[letter]; !ok {
			arr[letter] = true
		} else {
			return false
		}
	}
	return true
}

//Write a program to determine if a string contains only unique characters.
//Return true if it does and false otherwise.
//
//The string may contain any of the 128 ASCII characters. Characters are
//case-sensitive, e.g. 'a' and 'A' are considered different characters.
