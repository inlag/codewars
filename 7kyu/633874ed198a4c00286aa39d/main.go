package main

import "fmt"

func main() {
	fmt.Println(Game([4]int{2, 5, 8, 11}, [4]int{1, 4, 7, 10}, [4]int{0, 3, 6, 9}))
	fmt.Println(Game([4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}, [4]int{0, 9, 10, 11}))
	fmt.Println(Game([4]int{1, 6, 8, 9}, [4]int{0, 5, 10, 11}, [4]int{2, 3, 4, 7}))
}

func Game(frank, sam, tom [4]int) bool {
	var count = 0
	for i := 0; i < 4; i++ {
		if count >= 2 {
			return true
		}

		for ii := 0; ii < 4; ii++ {
			if frank[ii] > sam[i] && frank[ii] > tom[i] {
				count++
				break
			}
		}
	}

	return count >= 2
}
