<<COMMENT
Make, Copy & Remove

Learning Objectives:
* Escaping special characters in shell
* Removing, moving and making files
* Creating directories
* Basic shell scripting II
* Basic Conditionals I
* Googling

Theory
A great programmer is not someone who writes 100 code lines per minute, but someone who can work through technical details 
and understand what is important right now (and what is not) and use that info to achieve his objectives. 
This skill is not only vital for programmers, but IT technicians, system administrators, researchers etc.
You'll have most likely used Google in the previous task to learn about piping, but this task will require you to use it 
for all the commands, since we won't give out any hints on the commands used here.
Some characters in the shell are special, and will cause the shell to act in unexpected ways without another special character to cancel it's effect.
This task will also feature the first use of conditionals. 
You'll need to Google on how to write conditionals in the shell for this task, but we will write more about conditionals in the next task.

Instructions
These instructions seem daunting, but a wise friend called "Google" has offered to help you in these troubling times...
In the file files.sh, write the commands that do the following in the specified order:
1. Create files a, !, \ and "
2. Create a directory called `
3. Copy the file ! into `
4. If the environment variable called MOVE_A has the value yes, move the a file into the ` directory. If it's no, remove the file. Otherwise (variable is unset or is something else), leave the file as is.

Hints
You'll need to use Google for this task. Think about how you would phrase your
question, and then type it into Google. Modify that phrase until you get what
you want.
COMMENT

#!/bin/bash
touch a '!' '\' '"'
mkdir '`'
cp '!' '`'
if [ "$MOVE_A" = "yes" ]; then
	mv a '`'
elif [ "$MOVE_A" = "no" ]; then
	rm a
fi
