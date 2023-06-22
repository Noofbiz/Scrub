package scenes

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type MadeScene struct{}

func (*MadeScene) Type() string { return "Made in Engo!" }

func (s *MadeScene) Preload() {}

func (s *MadeScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	var animatable *common.Animationable
	var notanimatable *common.NotAnimationable
	var animSys = &common.AnimationSystem{}
	w.AddSystemInterface(animSys, animatable, notanimatable)

	var audioable *common.Audioable
	var notaudioable *common.NotAudioable
	w.AddSystemInterface(&common.AudioSystem{}, audioable, notaudioable)

	var mouseable *common.Mouseable
	var notmouseable *common.NotMouseable
	w.AddSystemInterface(&common.MouseSystem{}, mouseable, notmouseable)
}