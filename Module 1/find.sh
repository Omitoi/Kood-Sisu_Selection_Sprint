<<COMMENT
Find Your Way

Learning Objectives
* Regular expressions with wildcards
* Searching files
* Basic conditionals II

Instructions
Sometimes you need to look for something that you don't know if it even exists.
Create a file find.sh and in it, write a command which will look for and show in the current directory and its sub-folders:
* Everything that starts with an a, or
* All the files ending with a z, or
* All files starting with z and ending with a.

Hints
All you need to complete this task are man and find.
Use your knowledge from the previous task, and create your own tests here!
COMMENT

#!/bin/bash
find . \( -name 'a*' -o -name '*z' -o -name 'z*a' \)
