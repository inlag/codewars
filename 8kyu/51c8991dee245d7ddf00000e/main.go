package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.EqualFold(
			ReverseWords("hello world!"),
			"world! hello",
		),
	)
}

func ReverseWords(str string) string {
	strs := strings.Split(str, " ")

	var res string
	wordsCount := len(strs)
	for i := wordsCount - 1; i >= 0; i-- {
		res += strs[i] + " "
	}
	return res[:len(res)-1]
}

//Complete the solution so that it reverses all of the words within the string passed in.
//
//Example(Input --> Output):
//
//"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"
