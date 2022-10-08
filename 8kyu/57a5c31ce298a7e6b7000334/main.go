package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BinToDec("1001001"))
}

func BinToDec(bin string) int {
	res, _ := strconv.ParseInt(bin, 2, 64)
	return int(res)
}

//Complete the function which converts a binary number (given as a string) to a decimal number.
