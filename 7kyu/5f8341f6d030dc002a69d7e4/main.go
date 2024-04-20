package main

import (
	"fmt"
)

func main() {
	fmt.Println(LeastLarger([]int{4, 1, 3, 5, 6}, 0)) // 3
	fmt.Println(LeastLarger([]int{4, 1, 3, 5, 6}, 4)) // -1
	fmt.Println(LeastLarger([]int{1, 3, 5, 2, 4}, 0)) // 3

}

func LeastLarger(a []int, i int) (tt int) {
	tt = -1
	for index := range a {
		if index == i {
			continue
		}

		if a[index] > a[i] {
			if tt == -1 || tt != -1 && a[tt] > a[index] {
				tt = index
			}
		}
	}
	return
}
