package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := ToCsvText([][]int{
		{0, 1, 2, 3, 45},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34},
	})

	fmt.Println(
		strings.EqualFold(
			res,
			"0,1,2,3,45\n10,11,12,13,14\n20,21,22,23,24\n30,31,32,33,34",
		),
	)
}

func ToCsvText(array [][]int) string {
	var response string
	for _, ints := range array {
		for _, i2 := range ints {
			response += strconv.Itoa(i2) + ","
		}
		response = response[:len(response)-1] + "\n"
	}
	return response[:len(response)-1]
}
