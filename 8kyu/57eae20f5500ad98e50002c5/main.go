package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B"))
}

func NoSpace(word string) string {
	var b strings.Builder
	b.Grow(len(word))

	for _, w := range word {
		if !unicode.IsSpace(w) {
			b.WriteRune(w)
		}
	}
	return b.String()
}
