package main

import "fmt"

func main() {
	fmt.Println(OverTheRoad(3, 3))
	fmt.Println(OverTheRoad(2, 3))
	fmt.Println(OverTheRoad(7, 11))
	fmt.Println(OverTheRoad(23633656673, 310027696726))

}

func OverTheRoad(address int, n int) int {
	return 2*n + 1 - address
}
