package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f", Heron(1,8,8))
}

func Heron(a, b, c int) (area float32) {
	floatA,	floatB,	floatC := float64(a), float64(b), float64(c)
	semiPer := (floatA+floatB+floatC) / 2
	return float32(math.Sqrt(semiPer * (semiPer - floatA) * (semiPer - floatB) * (semiPer - floatC)))
}

