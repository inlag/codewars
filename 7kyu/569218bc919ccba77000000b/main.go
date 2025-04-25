package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(DateNbDays(4281, 5087, 2)) // "2024-07-03"
	fmt.Println(DateNbDays(4620, 5188, 2)) // "2021-09-19"
}

func DateNbDays(a0 int, a int, p int) string {
	var (
		a0F = float64(a0)
		aF  = float64(a)
		pF  = float64(p)
	)

	dateStr := "2016-01-01"
	if a0F > aF {
		return dateStr
	}
	startDate, _ := time.Parse(time.DateOnly, dateStr)

	var (
		dailyRate     = pF / 36000
		currentAmount = a0F
		days          = 0
	)

	for currentAmount < aF {
		currentAmount += currentAmount * dailyRate
		days++
	}

	resultDate := startDate.Add(time.Hour * 24 * time.Duration(days))
	return resultDate.Format("2006-01-02")
}
