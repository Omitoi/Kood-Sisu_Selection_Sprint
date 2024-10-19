package main

/*
Notes tool

Learning Objectives:
*	Participants will develop problem-solving skills as they design and implement a command-line tool that handles various user inputs, creates collections, and manages notes effectively.
	They will need to think logically to address different scenarios.
*	Participants will gain experience in validating user inputs, a crucial skill in software development.
	They will learn to ensure the program can handle unexpected user behavior, including handling invalid arguments or input.
*	This task offers an opportunity for participants to practice creating an interactive command-line interface.
	They will learn how to prompt users for input, provide options, and present results in a user-friendly and intuitive manner.
*	Participants will collaborate with team members to define responsibilities, discuss the program's structure, and integrate individual contributions into a cohesive solution.
	This teamwork aspect allows them to experience collaborative software development, a crucial skill in professional settings.

Instructions
You are tasked with creating a command-line tool that allows users to manage short single-line notes.
Using this application, users can create collections of notes, open and view them, add new notes, or remove existing notes.
The tool takes exactly one argument, which is the name of the collection.
To create or manage a collection coding_ideas the command would be:
$ ./notestool coding_ideas

Database
For each collection, a separate database is created.
The database is a plain text file with the same name as the collection, where notes are stored in separate rows.
If the collection does not exist, it should be created, if it does, it should be loaded.
Notes must persists between the runs of the tool.

Help message
If no argument is provided, the number of arguments is not one, or if the argument is help, the tool will display a brief help message that explains how to use the application.

$ ./notestool
Usage: ./todotool [TAG]

Tool functionality
If an argument is given, then the app should greet the user and display the tool menu.
From the menu the user must be able to do next operations:
1.	Display notes from the collection
2.	Add a note to the collection
3.	Remove a note from the collection
4.	Exit the program
	When user performs an operation, the user should be brought back to the menu, until he exits it from the menu.

You also have to make a brief description in Markdown, that at least:
*	Explains what the tool does
*	Explains the usage of the tool, with examples
*	Explains how the data is stored

Allowed packages
For this task you're allowed to use:
*	bufio
*	fmt
*	os
*	strconv
*	strings
*/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Note struct {
	Id         int
	NotRemoved bool
	Title      string
	Body       string
}

var args []string
var notesCollection string

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printText(input int) {
	switch input {
	case 0: //more than one argument
		fmt.Println("Can't read more than one notes set at a time, for more help, call the program with $ ./notestool help argument.")
	case 1: //help or no argument
		fmt.Println("\nWelcome to the notes tool!\n\nExample of use:\n$ ./notestool mynotes\n\nOpening the program requires an argument of the notes set you're planning to work with\nThen you follow the instructions and type in your notes\nTo select input, write your choice and press Enter")
	case 2: //welcome
		clearConsole()
		fmt.Println(string("\033[38;2;190;147;255m"), "\nWelcome to the notes tool!", string("\033[0m"))
	case 3: //main menu
		fmt.Println(string("\033[33m"), "\nNote collection:", string("\033[0m")+notesCollection)
		fmt.Println("\nSelect operation:\n1. \033[32mShow\033[0m notes.\n2. \033[34mAdd\033[0m a note.\n3. \033[31mDelete\033[0m a note.\n4. Exit.")
	case 4: //invalid Input
		clearConsole()
		fmt.Println("\n\033[31mInvalid input. Please follow the instructions\033[0m")
	case 5: //no Notes
		clearConsole()
		exec.Command("sleep", "0.5").Run()
		fmt.Println("\033[31mBly me!\033[0m There are no notes!")
		exec.Command("sleep", "0.5").Run()
	case 6: //exitting
		fmt.Println("\nExiting the notes. \033[38;2;190;147;255mGodspeed!")
	}

}

func easterEgg() {
	// Clear the screen
	clearConsole()

	// Define the maximum width and height for the terminal
	maxWidth := 25
	maxHeight := 25

	for i := 0; i < maxHeight; i++ {
		// Move the cursor to the current row and column
		fmt.Printf("\033[%d;%dH", i, i) // Move cursor

		// Print the moving text
		fmt.Print("\033[38;2;190;147;255mWeeee!\033[0m")

		// Sleep for a short duration to see the effect
		exec.Command("sleep", "0.05").Run()
		clearConsole()

		// Break if the text goes beyond the terminal width
		if i >= maxWidth || i >= maxHeight {
			break
		}
	}

	// Move cursor to the bottom of the terminal
	fmt.Print("\033[11;0H")
	clearConsole()
	printText(2)
}

func main() {

	//Read the arguments
	args = os.Args[1:]
	//If there are more than one argument
	if len(args) > 1 {
		printText(0) //more than one argument
	}
	//If there is no argument or "help" argument
	if len(args) == 0 || args[0] == "help" || args[0] == "Help" {
		printText(1) //help or no argument
		return
	}

	notesCollection = args[0]
	notes := readNotes(notesCollection) //load the notes

	reader := bufio.NewReader(os.Stdin)
	printText(2) //welcome
	for {
		notes = readNotes(notesCollection)
		printText(3) //main menu
		//Read the input cutting off possible whitespace around it
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		//Convert the input into int to condition and compare
		choice1, _ := strconv.Atoi(input1)
		if choice1 != 1 && choice1 != 2 && choice1 != 3 && choice1 != 4 && choice1 != 5 { //catching invalid input
			printText(4) //invalid Input
			continue
		}
		switch choice1 {
		//Showing notes
		case 1:
			if len(notes) == 0 {
				printText(5) //no Notes
				continue
			}
			clearConsole() //clear to print out notes
			fmt.Println("Notes: ")
			fmt.Println("Showing notes in " + notesCollection + "\n")
			noNotes := true
			for _, r := range notes {
				if r.NotRemoved {
					fmt.Printf("%03d %s - %s\n", r.Id, r.Title, r.Body)
					noNotes = false
				}
			}
			if noNotes {
				printText(5)
				continue
			}
			continue
		//Add a note
		case 2:
			fmt.Println("Title of the note: ")
			noteTitle, _ := reader.ReadString('\n')  //read the title
			noteTitle = strings.TrimSpace(noteTitle) //clean up of excess space
			if noteTitle == "" {
				printText(4) // invalid input
				continue
			}
			fmt.Println("Note body:")
			noteBody, _ := reader.ReadString('\n') // read the body
			noteBody = strings.TrimSpace(noteBody) //clean up of excess space
			if noteBody == "" {
				printText(4) // invalid input
				continue
			}
			nextIndex := len(notes) + 1 //creating index for the new note
			inputNote := Note{
				Id:         nextIndex,
				NotRemoved: true,
				Title:      noteTitle,
				Body:       noteBody,
			}
			newIndex := len(notes)
			notes = append(notes, inputNote)
			addNote(notesCollection, notes[newIndex])
			clearConsole()
			fmt.Println("\nNote saved")
			continue
		//Remove a note
		case 3:
			clearConsole() //clear to print out notes
			fmt.Println("Showing notes in " + notesCollection + "\n")
			noNotes := true
			for _, r := range notes {
				if r.NotRemoved {
					fmt.Printf("%03d %s - %s\n", r.Id, r.Title, r.Body)
					noNotes = false
				}
			}
			if noNotes {
				printText(5)
				continue
			}
			fmt.Println("\nEnter the id of note to remove or 0 to cancel:")
			idinput, _ := reader.ReadString('\n')
			idinput = strings.TrimSpace(idinput)
			theid, err := strconv.Atoi(idinput)
			println(theid)
			if theid == 0 { //detect cancel
				continue
			}

			if theid > len(notes) || err != nil {
				printText(4) //invalid Input
				continue
			} else {
				var noteToDelete Note
				for _, r := range notes { //find the note
					if r.Id == theid {
						noteToDelete = r
					}
				}
				if noteToDelete.NotRemoved { //check if it was deleted
					clearConsole()
					fmt.Printf("Removed note: %03d\n", theid)
					removeNote(notesCollection, theid) //remove the note
					continue
				}
				printText(4) //invalid Input
			}
		//Exit
		case 4:
			printText(6) //exitting
			return
		case 5:
			easterEgg()
			continue
		}
	}
}
