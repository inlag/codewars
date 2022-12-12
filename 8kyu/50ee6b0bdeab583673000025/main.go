package main

import "fmt"

func Namevar() string {
	var a string = "code"
	var b string = "wa.rs"
	var name string = a + b
	return name
}
func main() {
	fmt.Println(Namevar())
}
