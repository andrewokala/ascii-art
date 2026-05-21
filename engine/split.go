package engine

import "strings"

func SplitInput(input string) []string {
	result := strings.ReplaceAll(input, "\\n", "\n")
	return strings.Split(result, "\n")
}