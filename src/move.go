package main

import (
	"fmt"
	"errors"
	"regexp"
)

type Move struct {
	from [2]int
	to [2]int
}

func newMove(input string) (Move, error) {
	valid, err := regexp.MatchString("[A-H][1-8].[A-H][1-8]", input)

	if err != nil {
		return Move{}, err
	}

	if !valid {
		return Move{}, errors.New("Invalid format. [A-H][1-8].[A-H][1-8]")
	}

	fmt.Println("Success!", input, "is a valid move!")

	inp := []rune(input)

	from := [2]int{ 
		int(inp[0] - []rune("A")[0]), 
		int([]rune("8")[0] - inp[1]),
	}

	to := [2]int{
		int(inp[3] - []rune("A")[0]), 
		int([]rune("8")[0] - inp[4]),
	}

	fmt.Println(from, to)

	return Move{ from, to }, nil
}