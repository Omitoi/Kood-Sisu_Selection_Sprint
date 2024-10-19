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

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointDiff(p1, p2 Point) Point {
	if p1.X*p1.Y > p2.X*p2.Y {
		return p1
	} else {
		return p2
	}
}
