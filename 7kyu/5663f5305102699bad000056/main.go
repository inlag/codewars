package main

import (
	"fmt"
)

func main() {
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println(MxDifLg(s1, s2))

	s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 = []string{}
	fmt.Println(MxDifLg(s1, s2))

	s1 = []string{"pbdagezybuethyhtlnhhj", "rbfbmvheluvrqpqyxb", "tzvbjjsaveidllsdugsfx", "xnpzwzoqdehktjwqceh", "muhvjaozfhlckomnikwz", "oochblkrrdjeuxiqq", "pzlxcdqsbnlrceatfyn", "nzbrmdoannkuajiqbuj", "cnayujqcjibalfwsbvkj", "kfrhxsykvarjaparsmrqt", "pwqtzcwnejxwijvkjv", "nopthxmfffsqqzhrwkl", "lzqvwvpbpaxbtjogwudz"}
	s2 = []string{"lvwsdhhbomctwjersv", "icqueiyybumkooixvk", "eskszkkqdferyenz", "lbengwheunqkpkkaespcz", "xvhxhxkufaduvmujibe", "joqppwtqqouycklcb", "ayssjqppgsaymiqlikrra", "tgehdpngtrmwfdjoncct", "bfyutteqzsgiuorvnip", "dgonbhnuqvvfmvbcnaf", "driprqguyuqiovsy", "mictojodrbvhampmuq"}
	fmt.Println(MxDifLg(s1, s2))
}

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	var result int
	for _, value := range a1 {
		for _, value2 := range a2 {
			if x := abs(len(value) - len(value2)); x > result {
				result = x
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// package kata

// func MxDifLg(a1 []string, a2 []string) int {
//   	if len(a1) == 0 || len(a2) == 0 {
// 		return -1
// 	}

// var minLen, maxLen int

// 	if min(a1, true) > min(a2, true) {
// 		maxLen = min(a1, true)
// 		minLen = min(a2, false)
// 	} else {
// 		maxLen = min(a2, true)
// 		minLen = min(a1, false)
// 	}

// 	return maxLen - minLen
// }

// func min(a1 []string, y bool) (res int) {
// 	res = len(a1[0])
// 	for _, word := range a1 {
// 		lenWord := len(word)

// 		if y {
// 			if lenWord > res {
// 				res = lenWord
// 			}
// 			continue
// 		}

// 		if lenWord < res {
// 			res = lenWord
// 		}
// 	}
// 	return
// }
