package main

import "fmt"

func main() {
	// I smell a series!
	fmt.Println(Well([]string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}))

	// Fail!
	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "bad"}))

	// Publish!
	fmt.Println(Well([]string{"bad", "bad", "bad", "bad", "good", "bad"}))
}

func Well(x []string) string {
	var keys = make(map[string]uint)
	for _, val := range x {
		keys[val] += 1
	}

	val, ok := keys["good"]
	switch {
	case val <= 2 && ok:
		return "Publish!"
	case val > 2:
		return "I smell a series!"
	default:
		return "Fail!"

	}
}

//For every good kata idea there seem to be quite a few bad ones!
//
//In this kata you need to check the provided array (x) for good ideas 'good'
//and bad ideas 'bad'. If there are one or two good ideas, return 'Publish!', if
//there are more than 2 return 'I smell a series!'. If there are no good ideas,
//as is often the case, return 'Fail!'.
