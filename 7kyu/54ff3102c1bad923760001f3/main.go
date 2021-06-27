package main

import "fmt"

func main() {
	fmt.Println(GetCount(""))
}

func GetCount(str string) (count int) {
	for _, v := range str {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
