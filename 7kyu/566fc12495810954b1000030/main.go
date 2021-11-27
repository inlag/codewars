package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(NbDig(550, 5))
	fmt.Println(NbDig(5750, 0))
}
func NbDig(n int, d int) (out int) {
	for i := 0; i <= n; i++ {
		out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return
}
