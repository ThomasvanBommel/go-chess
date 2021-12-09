package main

import (
	"errors"
	"regexp"
	"strings"
)

type Move struct {
	from [2]int
	to [2]int
	distance [2]int
}

func parseMove(input string) [2]int {
	inp := []rune(input)

	return [2]int{ 
		int(inp[0] - []rune("A")[0]), 
		int([]rune("8")[0] - inp[1]),
	}
}

func stringToMove(input string) (*Move, error) {
	input = strings.ToUpper(input)
	valid, err := regexp.MatchString("[A-H][1-8].[A-H][1-8]", input)

	if err != nil { return &Move{}, err }
	if !valid { return &Move{}, errors.New("Invalid format: [A-H][1-8].[A-H][1-8]") }

	from := parseMove(input[0:2])
	to :=   parseMove(input[3:])

	return &Move{ from, to, Distance(from, to) }, nil
}