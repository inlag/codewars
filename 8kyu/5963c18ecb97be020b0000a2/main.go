package main

import "fmt"

func main() {
	fmt.Println(Derive(7, 8))   // "56x^7"
	fmt.Println(Derive(5, 9))   // "45x^8"
	fmt.Println(Derive(354, 2)) // "45x^8"
}

func Derive(coefficient, exponent int) string {
	var responseNumber = coefficient * exponent
	if exponent >= 2 {
		exponent = exponent - 1
	}

	return fmt.Sprintf("%vx^%v", responseNumber, exponent)
}

//This function takes two numbers as parameters, the first number being the
//coefficient, and the second number being the exponent.
//
//Your function should multiply the two numbers, and then subtract 1 from the
//exponent. Then, it has to print out an expression (like 28x^7). "^1" should
//not be truncated when exponent = 2.
//
//For example:
//
//derive(7, 8) In this case, the function should multiply 7 and 8, and then
//subtract 1 from 8. It should output "56x^7", the first number 56 being the
//product of the two numbers, and the second number being the exponent minus 1.
//
//derive(7, 8) --> this should output "56x^7"
//derive(5, 9) --> this should output "45x^8"
