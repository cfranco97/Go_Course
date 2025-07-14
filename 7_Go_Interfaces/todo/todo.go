package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Struct tags are made by using backticks - useful for metadata
type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo) // convert (public) struct data to json format
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) // will return the error coming from WriteFile in case something goes wrong!
}

// Constructor
func New(content string) (Todo, error) {
	//check for input error on creator, if so return an empty Note with errors!
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{

		Text: content,
	}, nil
}
