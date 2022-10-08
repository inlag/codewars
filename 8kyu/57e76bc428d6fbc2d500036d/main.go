package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(StringToArray("Robin Singh"))
}

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}
