package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Movie(1460, 8, 0.5))
	fmt.Println(Movie(0, 10, 0.95))
}

func Movie(card, ticket int, perc float64) int {
	var (
		a, count float64 = 0, 0
		b                = float64(card)
	)

	for math.Ceil(b) >= float64(a) {
		count++
		a += float64(ticket)
		b += float64(ticket) * math.Pow(perc, count)
	}
	return int(count)
}
