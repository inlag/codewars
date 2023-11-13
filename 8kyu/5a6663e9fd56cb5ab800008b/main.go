package main

import "fmt"

func main() {
	fmt.Println(CalculateYears(1))
	fmt.Println(CalculateYears(2))
	fmt.Println(CalculateYears(10))
}

func CalculateYears(years int) (result [3]int) {
	result = [3]int{years, 0, 0}
	for i := 1; i <= years; i++ {
		if i == 1 {
			result[1] += 15
			result[2] += 15
		}
		if i == 2 {
			result[1] += 9
			result[2] += 9
		}

		if i > 2 {
			result[1] += 4
			result[2] += 5
		}
	}
	return result
}
