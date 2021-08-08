package main

import "fmt"

func main() {
	fmt.Println(Divisors(10))
}

func Divisors(n int) int {
	var res int = 1
	for _, degree := range primeFactorization(n) {
		res *= degree + 1
	}
	return res
}

func primeFactorization(n int) map[int]int {
	var (
		result = make(map[int]int, 0)
		degree = 2
	)

	for n != 1 {
		t := n % degree
		if t == 0 {
			n = n / degree
			result[degree]++
		} else {
			degree++
		}
	}

	return result
}
