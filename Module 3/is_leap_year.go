package sprint

/*
Is Leap Year?

Learning Objectives:
*	Writing a function that returns a boolean value is a fundamental skill for implementing conditions in programming.
*	Creating a function for this specific task promotes modular programming, which is important for code organization and reusability.
*	Understanding the mathematical rules for leap years (divisibility by 4, except for certain conditions) requires applying mathematical logic in programming.
*	Determining whether a year is a leap year or not is a problem-solving task, and learning to tackle such problems is essential in software development.

Instructions
Write a Go function that takes an integer representing a year and returns a boolean value indicating whether that year is a leap year or not.
A leap year is a year that is evenly divisible by 4, except for years that are divisible by 100 but not by 400.
*/

func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
