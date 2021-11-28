package main

import "fmt"

func main() {
	fmt.Println(PrevMultOfThree(36))
	fmt.Println(PrevMultOfThree(952406))
}

func PrevMultOfThree(n int) interface{} {
	for {
		switch n % 3 {
		case 0:
			if n == 0 {
				return nil
			}
			return n
		}
		n = n / 10
	}
}
