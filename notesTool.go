package main

import (
	"bufio"
	"fmt"
	"os"
)

func deleteNote(notes []string, index int) []string {
	return append(notes[:index], notes[index+1:]...)
}

func main() {
	var notes []string
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
			fmt.Println("Notes:")
			for i, note := range notes {
				fmt.Printf("%d. %s\n", i+1, note)
			}
		case 2:
			var note string
			fmt.Print("Enter a note text: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				note = scanner.Text()
				notes = append(notes, note)
				fmt.Println("Note added successfully.")
			}
		case 3:
			var index int
			fmt.Print("Enter the index of the note to delete: ")
			fmt.Scan(&index)
			if index >= 1 && index <= len(notes) {
				notes = deleteNote(notes, index-1)
				fmt.Println("Note deleted successfully.")
			} else {
				fmt.Println("Invalid index.")
			}
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
