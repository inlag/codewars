package main

import "fmt"

func main() {
	fmt.Println(CountRedBeads(6))
}

func CountRedBeads(n int) int {
	if n <= 1 {
		return 0
	}

	return (n - 1) * 2
}
