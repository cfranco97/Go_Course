package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
	"example.com/go-project/todo"
)

// Interface: guarantees that a certain value has a certain method.
// Interfaces don't define the logic of the method, only that it exists.
// In this case whatever struct signs this interface - has to have a Save() method implemented in its code.
type saver interface {
	Save() error
}

// interfaces can also embed other interfaces
type outputtable interface {
	saver // will have the Save() method
	Display()
}

func main() {

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

// this function is just to exemplify how you can accept any type of value inside functions and use switch types.
// The app won't crash if you input a value type that is not on the switch!
func printSomething(value any) {

	typedVal, ok := value.(int)
	// check if the incoming value is in fact of type int, then make some calculations with it
	if ok {
		fmt.Println(typedVal + 1)
	}

	/* switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	} */
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// save any struct type data that signs the 'saver' interface
func saveData(data saver) error {

	err := data.Save()

	if err != nil {
		fmt.Print("Saving the data failed")
		return err
	}
	fmt.Print("Data saved successfully.")
	return nil
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
