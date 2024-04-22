package main

import (
	"fmt"
)

func main() {
	fmt.Println(Brightest([]string{"#108CF8", "#000000"}))
	fmt.Println(Brightest([]string{"#ABCDEF", "#123456"}))
	fmt.Println(Brightest([]string{"#00FF00", "#FFFF00"}))
	fmt.Println(Brightest([]string{"#FFFFFF", "#1234FF"}))
	fmt.Println(Brightest([]string{"#FFFFFF", "#123456", "#000000"}))
	fmt.Println(Brightest([]string{"#14ABFB", "#071F37", "#B01467", "#293530", "#2F3797", "#214DA2", "#D8E6D4", "#C1F811", "#BC21B5", "#61E2F5", "#AEA0B5"}))
	fmt.Println(Brightest([]string{"#E225FD", "#92E068", "#9B2796", "#EE6DBD", "#6F40DE", "#3E5A0D", "#F12999", "#A0AC72", "#16FCC8"}))
}
func Brightest(colors []string) string {
	vMax, index := "00", 0
	for i, color := range colors {
		v, r, g, b := "00", color[1:3], color[3:5], color[5:]
		if r > g && r > b {
			v = r
		} else if g > r && g > b {
			v = g
		} else {
			v = b
		}
		if v > vMax {
			vMax, index = v, i
		}
	}
	return colors[index]
}

// https://www.codewars.com/kata/62eb800ba29959001c07dfee
// One of the common ways of representing color is the RGB color model, in which the Red,
//  Green, and Blue primary colors of light are added together in various ways to reproduce
// a broad array of colors.

// One of the ways to determine brightness of a color is to find the value V of the
// alternative HSV (Hue, Saturation, Value) color model. Value is defined as the largest
// component of a color:

// V = max(R,G,B)
// You are given a list of colors in 6-digit hexidecimal notation #RRGGBB. Return the
//  brightest of these colors!

// For example,

// brightest(["#001000", "#000000"]) == "#001000"
// brightest(["#ABCDEF", "#123456"]) == "#ABCDEF"
// If there are multiple brightest colors, return the first one:

// brightest(["#00FF00", "#FFFF00", "#01130F"]) == "#00FF00"
// Note that both input and output should use upper case for characters A, B, C, D, E, F.
