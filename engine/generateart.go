package engine

func GenerateArt(lines []string, bannerMap map[rune][8]string) []string {
	var result []string

	for i, line := range lines {
		if line == "" {
			if i != 0 {
				result = append(result, "")
			}
			continue
		}

		rows := make([]string, 8)

		for _, ch := range line {
			block := bannerMap[ch]

			for r := 0; r < 8; r++ {
				rows[r] += block[r]
			}
		}
		result = append(result, rows...)
	}

	return result
}
