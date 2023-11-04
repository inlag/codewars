package main

import "fmt"

func main() {
	fmt.Println(AmIWilson(0))
	fmt.Println(AmIWilson(1))
	fmt.Println(AmIWilson(5))
	fmt.Println(AmIWilson(8))
	fmt.Println(AmIWilson(9))
}

func AmIWilson(n int) bool {
	if n == 5 ||
		n == 13 ||
		n == 563 {
		return true
	}

	return false
}
