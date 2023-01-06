package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(NumberToString(123))
	fmt.Println(NumberToString(999))
	fmt.Println(NumberToString(-100))
}

func NumberToString(n int) string {
	return strconv.Itoa(n)
}

//We need a function that can transform a number (integer) into a string.
//
//What ways of achieving this do you know?
//
//Examples (input --> output):
//123  --> "123"
//999  --> "999"
//-100 --> "-100"
