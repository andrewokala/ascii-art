package engine

import "fmt"

func Render(input []string) {
	for _, printable := range input {
		fmt.Println(printable)
	}
}
