package main

import "fmt"

func main() {
	fmt.Println(Fib(3))
	fmt.Println(Fib(5))
}

func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

//Create function fib that returns n'th element of Fibonacci sequence (classic
//programming task).
