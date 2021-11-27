package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Compose("abcd\nefgh\nijkl\nmnop", "qrst\nuvwx\nyz12\n3456"))
}
func Compose(s1, s2 string) string {
	var response string
	split1 := strings.Split(s1, "\n")
	split2 := strings.Split(s2, "\n")

	for i, ii := 0, len(split2)-1; i < len(split2); i++ {
		response += fmt.Sprintf(
			"%s\n",
			split1[i][:1+i]+split2[ii-i][:len(split2)-i],
		)
	}
	return response[:len(response)-1]
}
