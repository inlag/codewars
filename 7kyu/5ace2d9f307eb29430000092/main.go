package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ModifyMultiply("This is a string", 3, 5))
	fmt.Println(ModifyMultiply("dDRsZyqSZU koKyTkipcdMlnSdi xfZDa CYxCnMKF wLddTqqngmxtjEWQM izjvGnQBsnFOSJYyAKCRE YDHPrhTgOeGPLHHX zlBwsOkZJUkoMf JTjgywQjJSkyQQEeDeD gkDeiC kQthNBdP TT nsQ DVkGTbhPikkwUrZ xmlfRSyOLpKXOdnGgT WVuq JMkMwbbYeuvqorrvET PxyTwtPdUpqJrbO sKSjjN eJeNPxbQbNY ebDdBCDcFr AdkSqQfFHJVTlQRcHvIQS QwWAfQRaE NfMmY BjtrhlfUID h abHNXEijJduffChZzED PhognnbXqv ErZAauqTFm sTsVj hBLzRLkOSPdkukilSJumIZk JgpMWafMZhFgTVLAVjkp qUycHKlLRSDTod bHeFnvutLMzn uEySmoIPjyhVgbDEWMW epUxUaxexZhHeMU hFsHcokTKqkzAs Tb WaDnclPKMyUp brOLkHSzmLSgqiyUPSO", 38, 0))
}

func ModifyMultiply(str string, loc, num int) string {
	strs := strings.Split(str, " ")

	if num == 0 {
		num++
	}
	resArr := make([]string, num, num)
	for key := range resArr {
		resArr[key] = strs[loc]
	}

	return strings.Join(resArr, "-")
}

//You are to write a function that takes a string as its first parameter. This
//string will be a string of words.
//
//You are expected to then use the second parameter, which will be an integer,
//to find the corresponding word in the given string. The first word would be
//represented by 0.
//
//Once you have the located string you are finally going to multiply by it the
//third provided parameter, which will also be an integer. You are additionally
//required to add a hyphen in between each word.
//
//Example
//
//modifyMultiply ("This is a string",3,5)
