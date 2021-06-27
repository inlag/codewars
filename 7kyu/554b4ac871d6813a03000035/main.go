package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	var high, low int
	for i, v := range strings.Split(in, " ") {
		t, _ := strconv.Atoi(v)
		if i == 0 {
			high, low = t, t
		}

		if t > high {
			high = t
		}
		if t < low {
			low = t
		}
	}
	return fmt.Sprintf("%v %v", high, low)
}

func main() {
	fmt.Println(HighAndLow("-1 "))
}
