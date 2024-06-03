package main

import (
	"fmt"
	"strings"
	"unicode"
)

var keys = map[rune]rune{
	'G': 'A', 'D': 'E', 'R': 'Y', 'P': 'O',
	'L': 'U', 'K': 'I', 'A': 'G', 'E': 'D',
	'Y': 'R', 'O': 'P', 'U': 'L', 'I': 'K',
}

func Encode(str string) string {
	var response strings.Builder
	for _, val := range str {
		if keyValue, ok := keys[unicode.ToUpper(val)]; ok {
			if unicode.IsLower(val) {
				keyValue = unicode.ToLower(keyValue)
			}
			response.WriteRune(keyValue)
			continue
		}
		response.WriteRune(val)
	}
	return response.String()
}

func Decode(str string) string {
	return Encode(str)
}

func main() {
	fmt.Println(Encode("ABCD"))
	fmt.Println(Encode("GBCE"))
}

// https://www.codewars.com/kata/592a6ad46d6c5a62b600003f

// Introduction
// The GADERYPOLUKI is a simple substitution cypher used in scouting to encrypt messages. The encryption is based on short, easy to remember key. The key is written as paired letters, which are in the cipher simple replacement.

// The most frequently used key is "GA-DE-RY-PO-LU-KI".

//  G => A
//  g => a
//  a => g
//  A => G
//  D => E
//   etc.
// The letters, which are not on the list of substitutes, stays in the encrypted text without changes.

// Task
// Your task is to help scouts to encrypt and decrypt thier messages. Write the Encode and Decode functions.

// Input/Output
// The input string consists of lowercase and uperrcase characters and white . The substitution has to be case-sensitive.

// Example
//  Encode("ABCD")          // => GBCE
//  Encode("Ala has a cat") // => Gug hgs g cgt
//  Encode("gaderypoluki")  // => agedyropulik
//  Decode("Gug hgs g cgt") // => Ala has a cat
//  Decode("agedyropulik")  // => gaderypoluki
//  Decode("GBCE")          // => ABCD
