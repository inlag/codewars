package main

import "fmt"

func main() {
	fmt.Println(CloseCompare(3, 5, 3))       // => 0
	fmt.Println(CloseCompare(3, 5, 0))       // => -1
	fmt.Println(CloseCompare(6.0, 5.0, 0.0)) // => 1
	fmt.Println(CloseCompare(5.0, 5.0, 0.0)) // => 0
}

func CloseCompare(a, b, margin float64) int {
	if a-b > margin {
		return 1
	}
	if (-1)*(a-b) <= margin {
		return 0
	}
	return -1
}

//Create a function close_compare that accepts 3 parameters: a, b, and an
//optional margin. The function should return whether a is lower than, close to,
//or higher than b.
//
//a is considered "close to" b if margin is greater than or equal to the
//distance between a and b.
//
//Please note the following:
//
//When a is close to b, return 0.
//Otherwise...
//
//When a is less than b, return -1.
//
//When a is greater than b, return 1.
//
//If margin is not given, treat it as zero.
//
//Assume: margin >= 0
//
//Tip: Some languages have a way to make parameters optional.
//
//Example 1
//If a = 3, b = 5, and margin = 3, then close_compare(a, b, margin) should return 0.
//
//This is because a and b are no more than 3 numbers apart.
//
//Example 2
//If a = 3, b = 5, and margin = 0, then close_compare(a, b, margin) should return -1.
//
//This is because the distance between a and b is greater than 0, and a is less than b.
//
