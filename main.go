package main

import (
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/Scrub/scenes"
	"github.com/Noofbiz/Scrub/shaders"
)

func main() {
	common.AddShader(shaders.BubbleShader)
	engo.RegisterScene(&scenes.MadeScene{})
	engo.RegisterScene(&scenes.MainMenuScene{})
	engo.Run(engo.RunOptions{
		Title:                      "Scrub!",
		Width:                      640,
		Height:                     360,
		ScaleOnResize:              true,
		ApplicationMajorVersion:    0,
		ApplicationMinorVersion:    1,
		ApplicationRevisionVersion: 1,
	}, &scenes.SkeleScene{})
}
