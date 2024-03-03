package main

import (
	"fmt"
	"math"
)

func FindNextSquare(sq int64) int64 {
	result := math.Sqrt(float64(sq))
	if math.Mod(result, 1.0) != 0.0 {
		return -1
	}
	result++
	return int64(result * result)
}

func main() {
	fmt.Println(FindNextSquare(int64(121)))         // 144
	fmt.Println(FindNextSquare(int64(625)))         // 676
	fmt.Println(FindNextSquare(int64(319225)))      // 320356
	fmt.Println(FindNextSquare(int64(15241383936))) // 15241630849
	fmt.Println(FindNextSquare(int64(155)))         // -1
}

// https://www.codewars.com/kata/56269eb78ad2e4ced1000013
