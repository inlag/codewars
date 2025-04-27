package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}))
	fmt.Println(ContainAllRots("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy",
		"BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}))
	fmt.Println(ContainAllRots("", []string{}))
	fmt.Println(ContainAllRots("abc", []string{"abc", "abc", "abc"}))
}

func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}

	var count int
	for i := range arr {
		if strings.Contains(strng+strng+strng+strng, arr[i]) {
			count++
		}
	}
	return utf8.RuneCountInString(strng) == count
}
