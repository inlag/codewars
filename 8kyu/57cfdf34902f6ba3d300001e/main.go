package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(TwoSort([]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}))
}

func TwoSort(arr []string) string {
	sort.Strings(arr)

	var response string
	word := strings.Split(arr[0], "")
	for i := 0; i < len(word); i++ {
		response += word[i] + "***"
	}
	return response[:len(response)-3]
}

//You will be given a list of strings. You must sort it alphabetically
//(case-sensitive, and based on the ASCII values of the chars) and then return
//the first value.
//
//The returned value must be a string, and have "***" between each of its letters.
//
//You should not remove or add elements from/to the array.
