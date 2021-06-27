package main

import "fmt"

func main() {
	fmt.Println(Greet(""))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}
