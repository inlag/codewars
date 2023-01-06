package main

import "fmt"

func main() {
	fmt.Println(ReverseSeq(5))
}

func ReverseSeq(n int) []int {
	var res = make([]int, 0, n)
	for i := n; i > 0; i-- {
		res = append(res, i)
	}
	return res
}

//Build a function that returns an array of integers from n to 1 where n>0.
//
//Example : n=5 --> [5,4,3,2,1]
