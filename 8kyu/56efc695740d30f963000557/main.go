package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ToAlternatingCase(str string) string {
	var response strings.Builder
	for _, i2 := range str {
		if unicode.IsUpper(i2) {
			response.WriteString(strings.ToLower(string(i2)))
		} else {
			response.WriteString(strings.ToUpper(string(i2)))
		}
	}

	return response.String()
}

func main() {
	fmt.Println(ToAlternatingCase("HELLO WORLD"))
	fmt.Println(ToAlternatingCase("HELLO world"))
}
