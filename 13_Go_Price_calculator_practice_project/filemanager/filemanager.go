package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("Failed to open file!")
	}

	scanner := bufio.NewScanner(file)

	var lines []string //

	// read line by line - returns a bool = false if there are no more lines to scan off the file
	// while true - keep scanning and appending to lines slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check if scanner returned error
	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Could not read text from file!")
	}

	file.Close()
	return lines, nil
}
