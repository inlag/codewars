package main

import "fmt"

func main() {
	fmt.Println(Contamination("_3ebzgh4", "&")) //&&&&&&&&
}

func Contamination(text, char string) string {
	if char == "" {
		return ""
	}
	var response string
	for _ = range text {
		response += char
	}

	return response
}

//An AI has infected a text with a character!!
//
//This text is now fully mutated to this character.
//
//If the text or the character are empty, return an empty string.
//There will never be a case when both are empty as nothing is going on!!
//
//Note: The character is a string of length 1 or an empty string.
//
//Example
//text before = "abc"
//character   = "z"
//text after  = "zzz"
