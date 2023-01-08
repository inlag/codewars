package main

import "fmt"

func main() {
	fmt.Println(Rps("scissors", "paper"))
	fmt.Println(Rps("scissors", "rock"))
	fmt.Println(Rps("rock", "rock"))
}

func Rps(p1, p2 string) string {
	p1Type := getInst(p1)
	p2Type := getInst(p2)

	switch p1Type.Rps(p2Type) {
	case 0:
		return "Player 1 won!"
	case 1:
		return "Player 2 won!"
	default:
		return "Draw!"
	}
}

// 0 - win
// 1 - lose
// 2 - draw

func getInst(player string) value {
	switch player {
	case "scissors":
		return scissor(player)
	case "rock":
		return rock(player)
	default:
		return paper(player)
	}
}

type (
	scissor string
	rock    string
	paper   string
)

type value interface {
	Rps(sc value) uint8
	GetValue() string
}

func (s scissor) GetValue() string {
	return string(s)
}

func (s scissor) Rps(sc value) uint8 {
	switch sc.GetValue() {
	case "paper":
		return 0
	case "rock":
		return 1
	default:
		return 2
	}
}

func (r rock) Rps(sc value) uint8 {
	switch sc.GetValue() {
	case "scissors":
		return 0
	case "paper":
		return 1
	default:
		return 2
	}
}

func (s rock) GetValue() string {
	return string(s)
}

func (p paper) Rps(sc value) uint8 {
	switch sc.GetValue() {
	case "rock":
		return 0
	case "scissors":
		return 1
	default:
		return 2
	}
}

func (s paper) GetValue() string {
	return string(s)
}

//Rock Paper Scissors
//Let's play! You have to return which player won! In case of a draw return Draw!.
//
//Examples(Input1, Input2 --> Output):
//
//"scissors", "paper" --> "Player 1 won!"
//"scissors", "rock" --> "Player 2 won!"
//"paper", "paper" --> "Draw!"
