package main

import "fmt"

func main() {
	fmt.Println(Solve([]int{3, 4, 4, 3, 6, 3}))
}

func Solve(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		for ii := 0; ii < i; ii++ {
			if arr[i] == arr[ii] {
				arr = append(arr[:ii], arr[ii+1:]...)
				break
			}
		}
	}
	return arr
}
