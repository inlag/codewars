package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FakeBin("45385593107843568")) // 01011110001100111
}

func FakeBin(x string) string {
	var response = strings.Builder{}
	for _, val := range x {
		if val < 53 {
			response.WriteString("0")
		} else {
			response.WriteString("1")
		}
	}

	return response.String()
}

//Given a string of digits, you should replace any digit below 5 with '0' and
//any digit 5 and above with '1'. Return the resulting string.

//Note: input will never be an empty string
