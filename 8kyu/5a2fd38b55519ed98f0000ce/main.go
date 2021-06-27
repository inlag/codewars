package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold(MultiTable(5), "1 * 5 = 5\n2 * 5 = 10\n3 * 5 = 15\n4 * 5 = 20\n5 * 5 = 25\n6 * 5 = 30\n7 * 5 = 35\n8 * 5 = 40\n9 * 5 = 45\n10 * 5 = 50"))
}

func MultiTable(number int) string {
	var res string
	for i := 1; i < 10; i++ {
		res += fmt.Sprintf("%v * %v = %v\n", i, number, i*number)
	}
	res += fmt.Sprintf("10 * %v = %v", number, 10*number)
	return res
}
