package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Seven(477557102))
}

func Seven(n int64) []int {
	floatN := float64(n)
	count := math.Log10(floatN)

	for i := 0; i <= int(count); i++ {
		remN := int(floatN) % 10
		floatN = float64(int(floatN / 10) - (remN * 2))
		//floatN = math.Round(float64(int(floatN / 10) - (remN * 2)))
	}
	return []int{
		int(floatN),
		int(count),
	}
}
//func Seven(n int64) []int {
//	floatN := float64(n)
//	count := math.Log10(float64(floatN))
//
//	for i := 0; i < int(count)-1; i++ {
//		remN := int(floatN) % 10
//		floatN = floatN / 10
//		floatN = math.Round(float64(int(floatN) - (remN * 2)))
//	}
//	return []int{
//		int(floatN),
//		int(count - 1),
//	}
//}
