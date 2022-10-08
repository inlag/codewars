package main

import "fmt"

func main() {
	fmt.Println(Feast("great blue heron", "garlic naan"))
	fmt.Println(Feast("brown bear", "bear claw"))
}

func Feast(beast string, dish string) bool {
	if beast[0] == dish[0] {
		if beast[len(beast)-1] == dish[len(dish)-1] {
			return true
		}
		return false
	}
	return false
}
