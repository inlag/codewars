package main

import "fmt"

func main() {
	fmt.Println(MaxMultiple(2, 7))
	fmt.Println(MaxMultiple(3, 10))
	fmt.Println(MaxMultiple(7, 17))
	fmt.Println(MaxMultiple(10, 50))
	fmt.Println(MaxMultiple(37, 200))
	fmt.Println(MaxMultiple(7, 100))
}

func MaxMultiple(d, b int) int {
	for i := b; i > d; i-- {
		if i%d == 0 && i > 0 {
			return i
		}
	}
	return 0
}
