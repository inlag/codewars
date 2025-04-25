package main

import (
	"fmt"
	"sort"
)

func choreAssignment(chores []int) []int {
	sort.Ints(chores)
	workloads := make([]int, len(chores)/2)

	for i := 0; i < len(chores)/2; i++ {
		workloads[i] = chores[i] + chores[len(chores)-1-i]
	}

	sort.Ints(workloads)
	return workloads
}

func main() {
	chores1 := []int{1, 5, 2, 8, 4, 9, 6, 4, 2, 2, 2, 9}
	fmt.Println(choreAssignment(chores1)) // [7, 8, 8, 10, 10, 11]

	chores2 := []int{5, 2, 1, 6, 4, 4}
	fmt.Println(choreAssignment(chores2)) // [7, 7, 8]
}
