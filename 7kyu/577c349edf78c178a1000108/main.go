package main

import (
	"fmt"
)

func main() {
	for _, val := range XMasTree(5) {
		fmt.Println(val)
	}
}

func XMasTree(height int) []string {
	var response = make([]string, 0, height+2)

	for i := 0; i < height; i++ {
		starEndStr := getCopySymbols(height-i-2, "_")
		hashtag := getCopySymbols(i+i*1, "#")
		response = append(response, starEndStr+hashtag+starEndStr)
	}

	starEndStr := getCopySymbols(height-2, "_")
	response = append(response, starEndStr+"#"+starEndStr)
	response = append(response, starEndStr+"#"+starEndStr)

	return response
}

func getCopySymbols(count int, sym string) (res string) {
	for i := count; i >= 0; i-- {
		res += sym
	}
	return
}

// https://www.codewars.com/kata/577c349edf78c178a1000108
//
// Complete the function that returns a christmas tree of the given height. The height is passed
// through to the function and the function should return a list containing each line of the tree.

// XMasTree(5) should return : ['____#____', '___###___', '__#####__', '_#######_', '#########', '____#____', '____#____']
// XMasTree(3) should return : ['__#__', '_###_', '#####', '__#__', '__#__']
// Pad with underscores (_) so each line is the same length. Also remember the trunk/stem of the tree.

// Examples
// The final idea is for the tree to look like this if you decide to print each element of the list:

// n = 5 will result in:

// ____#____              1
// ___###___              2
// __#####__              3
// _#######_              4
// #########       -----> 5 - Height of Tree
// ____#____        1
// ____#____        2 - Trunk/Stem of Tree
// n = 3 will result in:

// __#__                  1
// _###_                  2
// #####          ----->  3 - Height of Tree
// __#__           1
// __#__           2 - Trunk/Stem of Tree
