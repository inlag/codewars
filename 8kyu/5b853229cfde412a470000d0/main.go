package main

import "fmt"

func main() {
	fmt.Println(TwiceAsOld(0, 0))
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	twiceSon := sonYearsOld * 2
	if twiceSon > dadYearsOld {
		return twiceSon - dadYearsOld
	}
	return dadYearsOld - twiceSon
}
