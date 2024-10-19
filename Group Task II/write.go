package main

import (
	"fmt"
	"os"
)

func addNote(databaseName string, note Note) error {
	// Look for file
	file, err := os.OpenFile(databaseName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// Write into file
	if _, err := file.WriteString(fmt.Sprintf("Id: %d\nNotRemoved: %t\nTitle: %s\nBody: %s\n", note.Id, note.NotRemoved, note.Title, note.Body)); err != nil {
		return err
	}
	return nil
}

func removeNote(databaseName string, noteId int) {
	// Read all notes from the file
	notes := readNotes(databaseName)

	// Mark the specified note as removed
	for i := range notes {
		if notes[i].Id == noteId {
			notes[i].NotRemoved = false
			break
		}
	}
	// Open the file for writing, truncating it first
	file, _ := os.OpenFile(databaseName, os.O_TRUNC|os.O_WRONLY, 0644)
	defer file.Close()
	
	// Write all notes back to the file
	for _, note := range notes {
		file.WriteString(fmt.Sprintf("Id: %d\nNotRemoved: %t\nTitle: %s\nBody: %s\n", note.Id, note.NotRemoved, note.Title, note.Body))
	}
}
