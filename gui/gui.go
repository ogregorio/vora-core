package gui

import (
	color "github.com/gookit/color"
)

var VERBOSE bool

const R = color.Red
const G = color.Green
const C = color.Cyan
const Y = color.Yellow
const P = color.White

//DMessage print a Debug Message to the console with given color "R(ed),G(reen),C(yan),Y(ellow),W(White)"
func DMessage(message string, color string) {
	if VERBOSE {
		switch color {
		case "R":
			R.Println(message)
		case "G":
			G.Println(message)
		case "C":
			C.Println(message)
		case "Y":
			Y.Println(message)
		case "P":
			P.Println(message)
		}
	}
}
