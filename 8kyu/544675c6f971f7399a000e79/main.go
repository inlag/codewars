package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func main() {
	fmt.Println(StringToNumber("12345"))
	fmt.Println(StringToNumber("-6"))
}
