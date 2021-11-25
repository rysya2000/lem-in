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

	file, err := os.Open("tests/" + fileName[1])
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	if len(text) == 0 {
		return nil, errors.New("ERROR: file is empty")
	}
	return text, nil
}

//FileProcessing gets required variables from array with input information
func FileProcessing(text []string) (Farm, string, string, int, error) {
	antFarm := &Farm{}
	var startRoom, endRoom, dot string
	var ant int
	var startCounter, endCounter int
	usedCoords := make(map[string]int)

	for i, t := range text {
		if i == 0 {
			temp, err := strconv.Atoi(t)
			if err != nil || temp <= 0 {
				return *antFarm, startRoom, endRoom, temp, errors.New("ERROR: invalid number of ants")
			}
			ant = temp
			continue
		}
		if t == "" {
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
		if len(t) != 0 && t[0] == '#' && !strings.Contains(t, " ") {
			continue
		}
		if strings.Contains(t, "-") {
			r := strings.Split(t, "-")
			err := antFarm.AddTunnels(r[0], r[1])
			if err != nil {
				return *antFarm, startRoom, endRoom, ant, err
			}
			continue
		}

		s := strings.Split(t, " ")
		if !strings.Contains(t, "-") && len(s) != 3 {
			return *antFarm, startRoom, endRoom, ant, errors.New("ERROR: invalid coordinates")
		}

		usedCoords[s[1]+s[2]]++
		if usedCoords[s[1]+s[2]] > 1 {
			return *antFarm, startRoom, endRoom, ant, errors.New("ERROR: invalid coordinates")
		}

		_, err1 := strconv.Atoi(s[1])
		_, err2 := strconv.Atoi(s[2])
		if err1 != nil || err2 != nil {
			return *antFarm, startRoom, endRoom, ant, errors.New("ERROR: invalid coordinates")
		}

		if dot == "s" {
			startRoom = s[0]
			dot = ""
			startCounter++
		}
		if dot == "e" {
			endRoom = s[0]
			dot = ""
			endCounter++
		}

		if s[0][0] == '#' || s[0][0] == 'L' {
			return *antFarm, startRoom, endRoom, ant, errors.New("ERROR: name start with L or #")
		}

		err := antFarm.AddRoom(s[0])
		if err != nil {
			return *antFarm, startRoom, endRoom, ant, err
		}
	}
	if startCounter != 1 {
		return *antFarm, startRoom, endRoom, ant, errors.New("ERROR: one start room only required")
	}

	if endCounter != 1 {
		return *antFarm, startRoom, endRoom, ant, errors.New("ERROR: one end room only required")
	}
	return *antFarm, startRoom, endRoom, ant, nil
}
