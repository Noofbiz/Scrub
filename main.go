package main

import (
	"github.com/EngoEngine/engo"

	"github.com/Noofbiz/Scrub/scenes"
)

func main() {
	engo.Run(engo.RunOptions{
		Title:         "Scrub!",
		Width:         640,
		Height:        360,
		ScaleOnResize: true,
	}, &scenes.MadeScene{})
}