package main

import "fmt"

func main() {
	iverson := NBAPlayer{Team: "76ers", Ppg: 11.2}
	jordan := NBAPlayer{Team: "bulls", Ppg: 20.2}
	fmt.Println(SumPpg(iverson, jordan))
}

type NBAPlayer struct {
	Team string
	Ppg  float64
}

func SumPpg(playerOne, playerTwo NBAPlayer) float64 {
	return playerOne.Ppg + playerTwo.Ppg
}
