package main

import "fmt"

func main() {
	fmt.Println(solution("a", "z"))
}

func solution(str, ending string) bool {
	var (
		countStr    = len(str)
		countEnding = len(ending)
	)

	if countEnding > countStr {
		return false
	}

	if str[countStr-countEnding:] == ending {
		return true
	}
	return false
}
