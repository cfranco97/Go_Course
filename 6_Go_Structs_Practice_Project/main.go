package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print("Note saved successfully.")
}

// will output note title + content (if everything worked)
func getNoteData() (string, string) {
	title := getUserInput("Note tittle:")

	content := getUserInput("note Content:")

	return title, content
}

// Read user input, takes in the prompt value.
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)  // creating a reader that listens on the cmdline (os.Standart input)
	text, err := reader.ReadString('\n') // where to stop reading the string? (at line break)

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // new text = text (but without the linebreak at the end.)
	text = strings.TrimSuffix(text, "\r") // new text = text (but without the return carriage at the end - windows problem.)
	return text
}
