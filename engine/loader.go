package engine

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(name string) (map[rune][8]string, error) {
	file := "banner/" + name + ".txt"

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading %v", file)
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	splitter := strings.Split(content, "\n")

	bannerMap := make(map[rune][8]string)

	ascii := 32
	for i := 1; i+8 < len(splitter); i += 9 {
		var block [8]string

		for j := 0; j < 8; j++ {
			block[j] = splitter[i+j]
		}
		bannerMap[rune(ascii)] = block
		ascii++
	}

	return bannerMap, nil
}
