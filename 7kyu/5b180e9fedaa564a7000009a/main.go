package main

import (
	"fmt"
	"strings"
	"unicode"
)

func solve(str string) string {
	uppers := 0

	for _, r := range str {
		if unicode.IsUpper(r) {
			uppers += 1
		}
	}

	if uppers > len(str)/2 {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}

func main() {
	fmt.Println(solve("coDE"))
	fmt.Println(solve("z"))
	fmt.Println(solve("A"))
	fmt.Println(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}
