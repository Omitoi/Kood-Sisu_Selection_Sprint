package sprint

/*
Casting

Learning Objectives:

*	Gain proficiency in type casting by successfully converting a float64 to an int, ensuring that fractional parts are rounded appropriately.
*	Develop an understanding of rounding techniques, including rounding to the nearest integer, and implement this logic accurately within the function.

Instructions
Write a function that accepts a float64 value, rounds it to the nearest integer, and then returns the result as an int.
*/

import (
	"math"
)

func Casting(n float32) int {
	return int(math.Round(float64(n)))
}
