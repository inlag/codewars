package main

import (
	"fmt"
	"strings"
)

func Calc(s string) int {
	str := fmt.Sprintf("%v", []byte("ABC"))
	ff := strings.Count(str, "7")
	return (ff * 7) - ff
}

func main() {
	fmt.Println(Calc("ABC"))
	fmt.Println(Calc("abcdef"))
}
