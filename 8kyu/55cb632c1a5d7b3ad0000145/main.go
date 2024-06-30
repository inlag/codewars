package main

import "fmt"

func main() {
	fmt.Println(HoopCount(19))
	fmt.Println(HoopCount(9))
}

func HoopCount(n int) string {
	if n < 10 {
		return "Keep at it until you get it"
	}
	return "Great, now move on to tricks"
}
