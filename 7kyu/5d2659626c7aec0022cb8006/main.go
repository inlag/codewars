package main

import (
	"fmt"
	"regexp"
)

func main() {
	var ch = make(chan int, 100)
	go BaumSweet(ch)
	for {
		fmt.Println(<-ch)
	}
}

func BaumSweet(ch chan<- int) {
	ch <- 1
	p := regexp.MustCompile(`(1|^)0(00)*(1|$)`)
	i := 1
	for {
		if p.MatchString(fmt.Sprintf("%b", i)) {
			ch <- 0
		} else {
			ch <- 1
		}
		i++
	}
}

// https://www.codewars.com/kata/5d2659626c7aec0022cb8006/train/go
//
// Wikipedia: The Baum–Sweet sequence is an infinite automatic sequence of 0s and 1s defined by the rule:

// bn = 1 if the binary representation of n contains no block of consecutive 0s of odd length;
// bn = 0 otherwise;

// for n ≥ 0.

// Define a generator function BaumSweet that sequentially generates the values of this sequence.

// It will be tested for up to 1 000 000 values.

// Note that the binary representation of 0 used here is not 0 but the empty string
// ( which does not contain any blocks of consecutive 0s ).
