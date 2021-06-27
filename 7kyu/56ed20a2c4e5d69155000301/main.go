package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Scale("abcd\nefgh\nijkl\nmnop", 2, 3))
	fmt.Println("aabbccdd\naabbccdd\naabbccdd\neeffgghh\neeffgghh\neeffgghh\niijjkkll\niijjkkll\niijjkkll\nmmnnoopp\nmmnnoopp\nmmnnoopp")
}

func Scale(s string, k, n int) string {
	if s == "" {
		return ""
	}

	var response = strings.Builder{}
	for _, word := range strings.Split(s, "\n") {
		var outWord string
		for _, letter := range word {
			for i := 0; i < k; i++ {
				outWord += string(letter)
			}
		}

		for i := 0; i < n; i++ {
			response.WriteString(outWord)
			response.WriteString("\n")
		}
	}
	return response.String()[:len(response.String())-1]
}
