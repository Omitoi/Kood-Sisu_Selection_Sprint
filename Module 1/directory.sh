<<COMMENT 
The Directory

Learning Objectives:

* Navigate directories.
* Read file content and output it.
* Printing the current directory you are in.

Theory
A lot of the programmer's tools are written with the command-line in mind. While some IDEs have integrated some of these tools, it's best to learn this basic skill early, so when you need it, you won't be completely lost.

Instructions:
* You'll need to have the unzip and the wget command installed.

Getting the package
You seem to be lost in a strange land called "The Shell", where the people speak in "commands". However, there is a package that could help you...
In the repository you cloned, download a zip file:
wget https://study.koodsisu.fi/resource/theDirectory.zip

Then, unzip the file with
unzip theDirectory.zip

Finally, change into the new directory with
cd theDirectory

Changing places
The package seems to contain instructions on how to navigate these lands, through so called README's. You'll need to figure out on how to read this.

Change directories until you find a file that starts with "You have found me". Then, output the contents of this file into the terminal (the one where you are writing commands in). 
Finally, you must print the working directory of the file.

Documenting your new knowledge
It may be wise to write down these new commands, so that others can learn as well...

Create a file called directory.sh. You must put the commands you used for outputting the file and printing the working directory in this file, separated by newlines. 
You must not put any command that outputs something else other than the correct file and working directory.

Git commit and push the directory.sh file. You need not commit the other files.
COMMENT

cd theDirectory
cd left
cd down
cd beginning
echo "You have found me! Now call a cat upon my file and print working directory."
echo "/app/directories/theDirectory/left/down/beginning"
