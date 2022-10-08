package main

import "fmt"

func main() {
	fmt.Println(HowMuchILoveYou(7))
	fmt.Println(HowMuchILoveYou(3))
	fmt.Println(HowMuchILoveYou(6))
	fmt.Println(HowMuchILoveYou(408))
}

func HowMuchILoveYou(i int) string {
	if i < len(ILoveYouArray) {
		return ILoveYouArray[i-1]
	}
	number := i % 6
	if number == 0 {
		return ILoveYouArray[5]
	}
	return ILoveYouArray[number-1]
}

var ILoveYouArray = []string{
	"I love you",
	"a little",
	"a lot",
	"passionately",
	"madly",
	"not at all",
}
