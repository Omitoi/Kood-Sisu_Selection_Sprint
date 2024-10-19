package sprint

/*
Season

Learning Objectives:
*	Develop the skill of using the switch statement in Go to efficiently handle multiple cases based on the value of an expression.
*	Practice returning informative error messages when dealing with invalid or unexpected input, enhancing the robustness of your code.
*	Learn to structure your code logically, making it more organized and readable, which is particularly relevant when working with multiple cases or conditions.

Instructions
While if or if-else statements are very useful, sometimes they are a little uncomfortable. switch statement is very useful if one needs to perform more checks for a value or if there are multiple cases for one check.
In this task you need to make a function that takes a string, that can contain a month name. If a month name is given, the season has to be returned. Otherwise return "invalid input: " with the input appended to it.

Months:
winter: jan, feb, dec
spring: mar, apr, may
summer: jun, jul, aug
autumn: sep, oct, nov
*/

func Season(month string) string {
	switch month {
	case "jan", "feb", "dec":
		return "winter"
	case "mar", "apr", "may":
		return "spring"
	case "jun", "jul", "aug":
		return "summer"
	case "sep", "oct", "nov":
		return "autumn"
	}
	return "invalid input: " + month
}
