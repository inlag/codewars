package main

import "fmt"

func main() {
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
}

func PrinterError(s string) string {
	var errors int
	for _, v := range s {
		switch v {
		case 'a', 'b', 'c', 'd', 'e', 'f',
			'g', 'h', 'i', 'j', 'k', 'l', 'm':
		default:
			errors++
		}
	}
	return fmt.Sprintf("%v/%v", errors, len(s))
}
