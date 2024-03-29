package scenes

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/Scrub/systems"
)

type SkeleScene struct{}

func (*SkeleScene) Type() string { return "Skeleboy Studios" }

func (s *SkeleScene) Preload() {
	engo.Files.Load("SkeleAnimation.png")
	engo.Files.Load("skele.wav")
	engo.Files.Load("p1.ttf")
}

func (s *SkeleScene) Setup(u engo.Updater) {
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

	w.AddSystem(&systems.TransitionSystem{Delay: 5, Next: "Made in Engo!"})

	var flashytextable *systems.FlashyTextAble
	w.AddSystemInterface(&systems.FlashyTextSystem{}, flashytextable, nil)

	ss := common.NewSpritesheetWithBorderFromFile("EngoAnimation.png", 640, 360, 1, 1)
	gopher := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
		common.AnimationComponent
	}{BasicEntity: ecs.NewBasic()}
	gopher.RenderComponent.Drawable = ss.Drawable(0)
	gopher.AnimationComponent = common.NewAnimationComponent(ss.Drawables(), 0.1)
	gopher.AnimationComponent.AddAnimation(&common.Animation{Name: "EngoAnimation", Frames: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, Loop: false})
	w.AddEntity(&gopher)
	gopher.AnimationComponent.SelectAnimationByName("EngoAnimation")

	bgm := struct {
		ecs.BasicEntity
		common.AudioComponent
	}{BasicEntity: ecs.NewBasic()}
	player, _ := common.LoadedPlayer("jingle.wav")
	bgm.AudioComponent = common.AudioComponent{Player: player}
	w.AddEntity(&bgm)
	player.Play()

	txt := struct {
		ecs.BasicEntity
		common.RenderComponent
		common.SpaceComponent
		systems.FlashyTextComponent
	}{BasicEntity: ecs.NewBasic()}
	fnt := &common.Font{
		URL:  "p1.ttf",
		FG:   color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		Size: 48,
	}
	fnt.CreatePreloaded()
	txt.Drawable = common.Text{
		Font: fnt,
		Text: "Made using Engo!",
	}
	txt.RenderComponent.Color = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	txt.Position = engo.Point{X: 100, Y: 250}
	txt.SetZIndex(2)
	txt.FlashyTextComponent.Times = []float32{0.15, 0.15, 0.15, 0.15}
	txt.FlashyTextComponent.TextColors = []color.NRGBA{
		{0x32, 0xCB, 0xFF, 0xFF},
		{0x00, 0xA5, 0xE0, 0xFF},
		{0x89, 0xA1, 0xEF, 0xFF},
		{0xEF, 0x9C, 0xDA, 0xFF},
	}
	w.AddEntity(&txt)
}
