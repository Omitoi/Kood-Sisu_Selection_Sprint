package sprint

/*
Circle Result

Learning Objectives:
*	Apply everything you learned from the previous tasks

Instructions
You'll need a new structure. This will represent a circle and various properties
of it.
The Circle structure should include the following fields in the exact order:

Radius - Float
Diameter - Float
Area - Float
Perimeter - Float

The function, called GetCircle, will take in the radius of the circle. It
should return a new Circle with all the fields filled with the correct values
from the radius given.
For π, you MUST use the value 3.14.
*/

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	res := Circle{
		Radius:    r,
		Diameter:  r * 2,
		Area:      3.14 * r * r,
		Perimeter: 2 * 3.14 * r,
	}
	return res
}
