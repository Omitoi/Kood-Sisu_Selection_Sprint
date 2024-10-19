package sprint

/*
Point Difference
Learning Objectives:
*	Read data and manipulate data from structures

Instructions
Using the same structure from the previous, given two Point structures, return
the one that is further ahead in the X/Y coordinates. If they are equal, return
either.
*/

import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {
	p.Text = ""
	p.Text += fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return p
}
