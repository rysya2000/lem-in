package lemin

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

//OpenFile opens, scans file and adds to array input information
func OpenFile(fileName []string) ([]string, error) {

	if len(fileName) < 2 {
		return nil, errors.New("please enter the filename")
	} else if len(fileName) > 2 {
		return nil, errors.New("too many arguments")
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

//FileProcessing gets required variables from array with input information
func FileProcessing(text []string) (Farm, string, string, int) {
	antFarm := &Farm{}
	var startRoom, endRoom, dot string
	var ant int

	for i, t := range text {
		if i == 0 {
			ant, _ = strconv.Atoi(t)
			continue
		}
		if t == "##start" {
			dot = "s"
			continue
		}
		if t == "##end" {
			dot = "e"
			continue
		}

		if strings.Contains(t, "-") {
			r := strings.Split(t, "-")
			antFarm.AddTunnels(r[0], r[1])
			continue
		}

		s := strings.Split(t, " ")
		if dot == "s" {
			startRoom = s[0]
			dot = ""
		}
		if dot == "e" {
			endRoom = s[0]
			dot = ""
		}

		antFarm.AddRoom(s[0])
	}
	return *antFarm, startRoom, endRoom, ant
}
