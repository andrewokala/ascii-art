package main

import (
	"fmt"
	"os"
	"ascii-art/engine"
)

func main() {
	input, banner, err := engine.Validator(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	bannerMap, err := engine.LoadBanner(banner)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := engine.SplitInput(input)
	art := engine.GenerateArt(lines, bannerMap)
	engine.Render(art)
}
