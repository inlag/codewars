package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Dative("ablak"))   // ablaknak
	fmt.Println(Dative("tükör"))   // tükörnek
	fmt.Println(Dative("keret"))   // keretnek
	fmt.Println(Dative("otthon"))  // otthonnak
	fmt.Println(Dative("virág"))   // virágnak
	fmt.Println(Dative("tett"))    // tettnek
	fmt.Println(Dative("rokkant")) // rokkantnak
	fmt.Println(Dative("rossz"))   // rossznak
	fmt.Println(Dative("őr"))      // rossznak
}

// Dative returns the dative case for a given Hungarian word.
func Dative(word string) string {
	var suffixNEK = "e, é, i, í, ö, ő, ü, ű"
	var suffixNAK = "a, á, o, ó, u, ú"

	for i := len(word) - 1; i >= 0; i-- {
		if strings.ContainsAny(word[i:], suffixNEK) {
			return word + "nek"
		}
		if strings.ContainsAny(word[i:], suffixNAK) {
			return word + "nak"
		}
	}

	return word

}

// https://www.codewars.com/kata/57fd696e26b06857eb0011e7
//
//
// Vowel harmony is a phenomenon in some languages. It means that
// "A vowel or vowels in a word are changed to sound the same (thus "in harmony.")"
// (wikipedia). This kata is based on vowel harmony in Hungarian.

// Task:
// Your goal is to create a function dative() (Dative() in C#) which returns the
// valid form of a valid Hungarian word w in dative case i. e. append the correct
// suffix nek or nak to the word w based on vowel harmony rules.

// Vowel Harmony Rules (simplified)
// When the last vowel in the word is

// a front vowel (e, é, i, í, ö, ő, ü, ű) the suffix is -nek
// a back vowel (a, á, o, ó, u, ú) the suffix is -nak
// Examples:
// Dative("ablak") ==> "ablaknak", nil
// Dative("szék") ==> "széknek", nil
// Dative("otthon") ==> "otthonnak", nil
// Preconditions:
// To keep it simple: All words end with a consonant :)
// All strings are unicode strings.
// There are no grammatical exceptions in the tests.
