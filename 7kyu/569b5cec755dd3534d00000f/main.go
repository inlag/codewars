package main

import "fmt"

func main() {
	fmt.Println(NewAvg([]float64{1400, 30001, 6, 7, 9, 11, 15, 121}, 2000))
}

func NewAvg(arr []float64, navg float64) int64 {
	var sum float64
	for _, element := range arr {
		sum += element
	}
	result := int64(int64(navg)*int64(len(arr)+1) - int64(sum))

	if result < 0 {
		return -1
	}
	return result
}
