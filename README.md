## Code Structure

### Main Components

- **`deleteNote` Function:**
  - Deletes a note from the notes slice.
  - Takes a slice of notes and an index as arguments.
  - Returns the modified notes slice.

- **`loadNotes` Function:**
  - Loads notes from a file specified by the collection name.
  - Checks if the file exists; if not, creates a new file.
  - Reads the file and returns a slice of notes.

- **`saveNotes` Function:**
  - Saves notes to the file specified by the collection name.
  - Takes the collection name and a slice of notes as arguments.
  - Writes the notes to the file.

- **`main` Function:**
  - Validates the command-line arguments; exits if the collection name is missing.
  - Loads existing notes from the specified collection file.
  - Displays a menu for operations: Show notes, Add note, Delete note, Exit.
  - Performs the selected operation based on user input.

### User Interface

The user interacts with the tool through a simple command-line interface. The tool displays a menu with options to view, add, or delete notes. The user enters the corresponding numeric choice to perform the desired operation.

## Instructions

1. **Show Notes:**
   - Displays the existing notes in the collection with their respective indices.

2. **Add Note:**
   - Prompts the user to enter a new note and adds it to the collection.

3. **Delete Note:**
   - Asks the user for the index of the note to delete and removes it from the collection.

4. **Exit:**
   - Saves the current notes to the collection file and exits the tool.

## Error Handling

- If the specified collection file does not exist, it is created.
- Invalid user inputs are handled with appropriate error messages.

## Running the Code

Ensure you are in the correct directory and have the necessary permissions:

```sh
$ go build notestool.go
$ ./notestool yourlibraryname