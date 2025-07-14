package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct tags are made by using backticks - useful for metadata
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Display Note data function.
// (note Note) is a receiver for this function,in this case it will work with the used structure 'Note'
func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") // replace all 'space characters' from note.title to 'underscores'
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note) // convert (public) struct data to json format
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644) // will return the error coming from WriteFile in case something goes wrong!
}

// Constructor - use *Note if you want to avoid copies, but working with simple structs, its fine to avoid the extra overhead when working with pointers(only used on functions that mutate/edit the structure itself).
func New(title, content string) (Note, error) {

	//check for input error on creator, if so return an empty Note with errors!
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
