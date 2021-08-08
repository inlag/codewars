package main

import "fmt"

func main() {
	fmt.Println(PartList([]string{"I", "wish", "I", "hadn't", "come"}))
}

func PartList(arr []string) string {
	var resultStr string

	for i := 0; i < len(arr)-1; i++ {
		var str string
		for i2, word := range arr {
			str += word

			if i == i2 {
				str += ","
			}

			str += " "
		}
		resultStr += "(" + str[:len(str)-1] + ")"
	}
	return resultStr
}
