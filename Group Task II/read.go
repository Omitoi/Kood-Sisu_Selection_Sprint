package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func readNotes(databaseName string) []Note {
	// Find the file
	file, err := os.Open(databaseName)
	if err != nil {
		return []Note{}
	}
	defer file.Close()

	// Define variable to be used
	var notes []Note
	scanner := bufio.NewScanner(file)

	var id int
	var notRemoved bool
	var title, body string

	// Go through each line of the file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trim whitespace to avoid leading/trailing spaces

		switch {
		case strings.HasPrefix(line, "Id: "):
			id, err = strconv.Atoi(strings.TrimPrefix(line, "Id: "))
			if err != nil {
				// Handle the conversion error (e.g., log it or skip the note)
				continue
			}

		case strings.HasPrefix(line, "NotRemoved: "):
			notRemoved, err = strconv.ParseBool(strings.TrimPrefix(line, "NotRemoved: "))
			if err != nil {
				// Handle the conversion error
				continue
			}

		case strings.HasPrefix(line, "Title: "):
			title = strings.TrimPrefix(line, "Title: ")

		case strings.HasPrefix(line, "Body: "):
			body = strings.TrimPrefix(line, "Body: ")

			// After reading the body, construct the Note and append to the list
			note := Note{
				Id:         id,
				NotRemoved: notRemoved,
				Title:      title,
				Body:       body,
			}
			notes = append(notes, note)
		}
	}

	if err := scanner.Err(); err != nil {
		// Handle any errors that occurred during the scan
	}

	return notes
}
