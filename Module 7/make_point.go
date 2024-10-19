package sprint

/*
New Struct

Learning Objectives:
*	Understand how to compose data structures
*	Creating data structures (constructors)

Instructions
You'll need to compose a structure called Point, which will contain the
following fields (in the exact order):

X - float
Y - float
Text - string

This structure will represent text at a specific point (for example, on a graph
or video game).
You will need to write a function that will return a newly created struct of
this with fields filled from the parameters provided.
*/

type Point struct {
	X    float32
	Y    float32
	Text string
}

func MakePoint(x, y float32, text string) Point {
	res := Point{
		X:    x,
		Y:    y,
		Text: text,
	}
	return res
}
