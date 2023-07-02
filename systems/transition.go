package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type TransitionSystem struct {
	Next    string
	Delay   float32
	elapsed float32
}

func (s *TransitionSystem) Remove(basic ecs.BasicEntity) {}

func (s *TransitionSystem) Update(dt float32) {
	s.elapsed += dt
	if s.elapsed > s.Delay || engo.Input.Mouse.Action == engo.Press {
		engo.SetSceneByName(s.Next, true)
	}
}
