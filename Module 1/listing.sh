<<COMMENT
Command Line Mastery: The Listing

Learning Objectives:

 *Grasping advanced command line flags.
 *Mastering directory listing with custom conditions.
 *Sorting files based on access times.

Instructions:
1. List Like a Pro
You've got the basics down, but now it's time to refine your command line prowess. Your task: Display directory contents with a bit of flair.

2. Your Mission
In a file named listing.sh, craft a command line that:
* Lists all non-hidden files and directories in the current directory.
* Omits the entries . and ...
* Uses only commas and spaces as separators.
* Sorts results by access time, with the most recent first.
* Appends a / to directory names.

Hint
For the expert touch, remember: the man command holds many secrets. Seek its wisdom.
COMMENT
ls -p -m -u
