package main

import "fmt"

func main() {
	fmt.Println(Gps(
		12,
		[]float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25},
	))
}
func Gps(s int, x []float64) int {
	var (
		maxValue float64
		floatS   = float64(s)
	)
	for i := 0; i < len(x)-1; i++ {
		deltaDistance := (x[i] - x[i+1]) * -1
		value := (3600 * deltaDistance) / floatS
		if value > maxValue {
			maxValue = value
		}
	}
	return int(maxValue)
}
