<<COMMENT
Needles In a Haystack

Learning Objectives:
* Learn to search for given lines in files
* Basic regular expressions
* Basic piping

Theory
grep is among the most useful commands in Unix shells. It allows you find where specific text occurs at specific lines. 
This can be helpful in, for example, reading logs.

Another important concept in the shell world is piping. Piping allows you to put the output of one command as the input of the next command.
You'll need to Google about piping for this task.

Instructions
Download this file that you can test your commands against:

wget https://study.koodsisu.fi/resource/hay.txt

Finding the needles...
This file seems to contain a lot of repeated lines, but something in you says that some lines are not like the others.

You must find the lines that contain the word 'not'.

...and counting them
That's a lot of needles, now if there was some way to count them accurately...
Your script must output the number of lines that contain the word 'not'

Hints
The only commands you need are grep and wc. And of course, there's the always helpful man command...
COMMENT

grep -w 'not' hay.txt | wc -l
