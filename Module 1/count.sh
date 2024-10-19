<<COMMENT
Count Multiply

Learning Objectives:
* Basic variables
* Bash multiplication operation
* Basic text formatting
* Basic shell scripting I

Theory
Variables are kinda like containers, they store some data for later usage. 
For example, a variable called "numberOfLives" in a game could be used to show the player how many lives he has when required. 
Variables can also changed, for example if the player loses or gains a live.
We can also do mathematical operations on these variables. 
Usually in programming the basic operations we can do are adding/subtracting, multiplication, dividing and modulo (return the remainder after division, if any).
We can format variables into text, and print it in various ways, like using vertical tabs and newlines to affect the placement of the text.

Instructions
It's time to see your new skills in style!

* Find the number of files/directories in the current directory and sub-directories (and their sub-directories and so on).
* Multiply that number by 5.
* Print the number with a horizontal and vertical tab at the beginning, and a vertical tab and newline at the end (see usage). The text before the number should be the same as in usage.

Usage
$ ./count.sh 

    Total files * 5: 25

$ ./count.sh | cat -e
        ^KTotal files * 5: 25^K$
$

Hints
Look into the command called printf
COMMENT

#!/bin/bash
count=$(find . -type f -o -type d | wc -l)
result=$((count * 5))
printf "\t\vTotal files * 5: $result\v\n"
