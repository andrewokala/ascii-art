package engine

import (
	"fmt"
)

func Validator(argument []string) (string, string, error) {
	if len(argument) < 2 || len(argument) > 3 {
		return "", "", fmt.Errorf("invalid number of arguments")
	}

	input := argument[1]
	banner := "standard"

	if len(argument) == 3 {
		banner = argument[2]
	}

	valid := map[string]bool {
		"standard": 	true,
		"shadow": 		true,
		"thinkertoy":	true,
	}

	if !valid[banner] {
		return "", "", fmt.Errorf("invalid banner")
	}

	for i := 0; i < len(input); i++ {
		if input[i] == '\n'{
			continue
		}

		if input[i] < 32 || input[i] > 126 {
			return "", "", fmt.Errorf("character invalid")
		}
	}

	return input, banner, nil
}
