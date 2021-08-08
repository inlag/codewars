package main

import "fmt"

func main() {
	fmt.Println(WrapPresent(13, 13, 13))
}

func WrapPresent(height, width, length int) int {
	if height <= width && height <= length {
		return height * 4 + width * 2 + length * 2 + 20
	} else if width <= length {
		return height * 2 + width * 4 + length * 2 + 20
	}

	return height * 2 + width * 2 + length * 4 + 20
}
