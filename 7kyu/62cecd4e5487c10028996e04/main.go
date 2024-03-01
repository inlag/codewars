package main

import "fmt"

func main() {
	fmt.Println(RacePodium(6))
	fmt.Println(RacePodium(11))
	fmt.Println(RacePodium(100000))
	fmt.Println(RacePodium(2430))
	fmt.Println(RacePodium(7))
}

func RacePodium(blocks int) [3]int {
	if blocks == 7 {
		return [3]int{2, 4, 1}
	}
	x := (blocks + 5) / 3
	return [3]int{x - 1, x, blocks - 2*x + 1}
}

// https://www.codewars.com/kata/62cecd4e5487c10028996e04
//
//The national go-kart racing competition is taking place at your local town and
//you've been called for building the winners podium with the available wooden
//blocks. Thankfully you are in a wood-rich area, number of blocks are always at
//least 6.
//
//Remember a classic racing podium has three platforms for first, second and
//third place. First place is the highest and second place is higher than third.
//Also notice that platforms are arranged as 2nd - 1st - 3rd.
//
//The organizers want a podium that satisfies:
//
//The first place platform has the minimum height possible
//The second place
//platform has the closest height to first place
//All platforms have heights
//greater than zero. Task Given the numbers of blocks available, return an array
/// tuple or another data structure depending on the language (refer sample
//tests) with the heights of 2nd, 1st, 3rd places platforms.
//
//Examples (input -> output)
//11 ->   [4, 5, 2]
//6  ->   [2, 3, 1]
//10 ->   [4, 5, 1]
