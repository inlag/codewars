package main

import "fmt"

func main() {
	fmt.Println(IsTriangle(4, 2, 3))
}
func IsTriangle(a, b, c int) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if (a+b > c) &&
		(a+c > b) &&
		(c+b > a) {
		return true
	} else {
		return false
	}
}
