package main

import "fmt"

func main() {
	fmt.Println(CountSheeps(
		[]bool{
			true, true, true, false,
			true, true, true, true,
			true, false, true, false,
			true, false, false, true,
			true, true, true, true,
			false, false, true, true,
		},
	))
}

func CountSheeps(numbers []bool) int {
	var count int = len(numbers)
	for _, num := range numbers {
		if !num {
			count--
		}
	}
	return count
}
