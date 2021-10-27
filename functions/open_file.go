package lemin

import (
	"bufio"
	"errors"
	"os"
)

func OpenFile(fileName []string) ([]string, error) {

	if len(fileName) < 2 {
		return nil, errors.New("Please enter the filename!")
	} else if len(fileName) > 2 {
		return nil, errors.New("Too many arguments")
	}

	file, err := os.Open(fileName[1])
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text, nil
}
