package main

import "fmt"

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
}

func NbYear(p0 int, percent float64, aug int, p int) (result int) {
	for {
		if p0 >= p {
			return
		}
		p0 = p0 + aug + int(float64(p0)*(percent/100))
		result++
	}
}
