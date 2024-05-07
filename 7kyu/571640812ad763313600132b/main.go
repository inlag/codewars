package main

import "fmt"

func main() {
	fmt.Println(AlexMistakes(10, 120))
	fmt.Println(AlexMistakes(11, 120))
	fmt.Println(AlexMistakes(9, 180))
	fmt.Println(AlexMistakes(20, 120))
	fmt.Println(AlexMistakes(20, 125))
}

func AlexMistakes(numberOfKatas, timeLimit int) int {
	if numberOfKatas*6 >= timeLimit {
		return 0
	}

	timeLimit -= numberOfKatas * 6
	var (
		res      = 0
		duration = 5
	)
	for timeLimit >= duration {
		timeLimit -= duration
		duration *= 2
		res++
	}
	return res
}

// https://www.codewars.com/kata/571640812ad763313600132b
//
// Alex is transitioning from website design to coding and wants to sharpen his skills with CodeWars.
// He can do ten kata in an hour, but when he makes a mistake, he must do pushups.
// These pushups really tire poor Alex out, so every time he does them they take twice as long.
// His first set of redemption pushups takes 5 minutes. Create a function, alexMistakes, that takes
// two arguments: the number of kata he needs to complete, and the time in minutes he has to complete them.
// Your function should return how many mistakes Alex can afford to make.
