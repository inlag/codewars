package main

import "fmt"

func main() {
	fmt.Println(Fusc(19))
}

func Fusc(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if n%2 == 0 {
		return Fusc(n / 2)
	}
	return Fusc((n-1)/2) + Fusc((n+1)/2)
}
