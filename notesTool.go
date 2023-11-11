package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	toEncrypt, encoding, message := getInput()
	if encoding == "1" {
		fmt.Println(message, toEncrypt)
		// Command to call the encryption/decription function 1 and output the final code
	} else if encoding == "2" {
		fmt.Println(message)
	} else if encoding == "3" {
		fmt.Println(message)
	} else if encoding == "4" {
		fmt.Println(message)
	} /* else {
		fmt.Println(customEncDec(message, toEncrypt))
	}*/
	// Command to call the encryption/decription function 3 and output the final code
}

// INPUT AND OUTPUT
func getInput() (toEncrypt bool, encoding string, message string) {
	fmt.Println("Welcome to the notes tool!")
	// Command for output of the welcome phrase
	fmt.Println("\n Select operation (1/4):\n 1. Show notes. \n 2. Add a note. \n 3. Delete a note. \n 4. Exit")
	var SelectOp int
	fmt.Scan(&SelectOp)
	// what to do with notes
	if SelectOp == 1 {
		getInput()
		//Show notes
	} else if SelectOp == 2 {
		getInput()
		// add a note
	} else if SelectOp == 3 {
		getInput()
		//deleteing a note
	} else if SelectOp == 4 {
		//Exit
	} else {
		fmt.Println("Incorrect input, please try again")
		getInput()
	}
	// Logic by which the toEncrypt variable is assigned true or false, also if the user enters incorrect data, the program is started again
	/*fmt.Println("Select cypher (1/2/3):\n 1. ROT13. \n 2. Reverse \n 3. CustomCipher")
	fmt.Scan(&encoding)
	if encoding == "1" || encoding == "2" || encoding == "3" {
		fmt.Println("Enter the message:")
		message = getMessage()
		// function call for text input
	} else {
		fmt.Println("Incorrect input, please try again")
		getInput()
	}*/
	return toEncrypt, encoding, message
	// Returns the entered data to the main function
}

func getMessage() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
	// Used the bafio library and os to enter the code with spaces, since fmt.Scan does not allow for this
}
