package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Accum("atr"))
}

func Accum(s string) string {
	var result string
	for i, v := range s {
		result += strings.ToUpper(string(v))
		for ii := 0; ii < i; ii++ {
			result += strings.ToLower(string(v))
		}
		result += "-"
	}
	return result[:len(result)-1]
}
