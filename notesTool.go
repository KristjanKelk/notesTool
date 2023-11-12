package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func deleteNote(notes []string, index int) []string {
	return append(notes[:index], notes[index+1:]...)
}

func loadNotes(collection string) []string {
	// check if file exists
	_, err := os.Stat(collection)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(collection)
		if err != nil {
			fmt.Println("Failed to create collection")
			return []string{}
		}
		file.Close()
	}

	data, err := ioutil.ReadFile(collection)
	if err != nil {
		return []string{}
	}
	notes := strings.Split(string(data), "\n")
	var filteredNotes []string
	for _, note := range notes {
		if note != "" {
			filteredNotes = append(filteredNotes, note)
		}
	}

	return filteredNotes
}

func saveNotes(collection string, notes []string) {
	ioutil.WriteFile(collection, []byte(strings.Join(notes, "\n")), 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a collection name.")
		os.Exit(1)
	}

	collection := os.Args[1]
	notes := loadNotes(collection)

	fmt.Println("Welcome to notes tool. \nPlease enter only numerical values, example: 1")
	for {

		fmt.Println("\n1. Show notes")
		fmt.Println("2. Add note")
		fmt.Println("3. Delete note")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Select operation: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("")
			fmt.Println("Notes:")
			for i, note := range notes {
				fmt.Printf("%03d - %s\n", i+1, note)
			}
		case 2:
			var note string
			fmt.Print("")
			fmt.Print("Enter a note text: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				note = scanner.Text()
				notes = append(notes, note)
				fmt.Println("Note added successfully.")
			}
		case 3:
			var index int
			fmt.Println("")
			fmt.Print("Enter the index of the note to delete: ")
			fmt.Scan(&index)
			if index >= 1 && index <= len(notes) {
				notes = deleteNote(notes, index-1)
				fmt.Println("Note deleted successfully.")
			} else {
				fmt.Println("Invalid index.")
			}
		case 4:
			saveNotes(collection, notes)
			os.Exit(0)
		default:
			fmt.Println("")
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}