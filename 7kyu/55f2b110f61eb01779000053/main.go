package main

import "fmt"

func main() {
	fmt.Println(GetSum(1, 1))      // 1
	fmt.Println(GetSum(1, 2))      // 3
	fmt.Println(GetSum(321, 123))  // 44178
	fmt.Println(GetSum(-321, 123)) // -44055
}
func GetSum(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return a + GetSum(a-1, b)
	}
	return a + GetSum(a+1, b)
}

//Given two integers a and b, which can be positive or negative, find the sum of
//all the integers between and including them and return it. If the two numbers
//are equal return a or b.
//
//Note: a and b are not ordered!
//
//Examples (a, b) --> output (explanation)
//(1, 0) --> 1 (1 + 0 = 1)
//(1, 2) --> 3 (1 + 2 = 3)
//(0, 1) --> 1 (0 + 1 = 1)
//(1, 1) --> 1 (1 since both are same)
//(-1, 0) --> -1 (-1 + 0 = -1)
//(-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)
//Your function should only return a number, not the explanation about how you get that number.
