package scenes

import (
	"bytes"
	"image/color"
	"strconv"

	"golang.org/x/image/font/gofont/gomono"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/Noofbiz/Scrub/shaders"
	"github.com/Noofbiz/Scrub/systems"
	"github.com/Noofbiz/pixelshader"
)

type MainMenuScene struct{}

func (s *MainMenuScene) Type() string { return "Main Menu" }

func (s *MainMenuScene) Preload() {
	engo.Files.Load("pleasantcreek.wav")
	engo.Files.Load("ui/BubblegumSans-Regular.ttf")
	engo.Files.Load("ui/menu_card.png")
	engo.Files.Load("ui/green_button00.png")
	engo.Files.Load("ui/yellow_button00.png")
	engo.Files.Load("ui/red_boxCross.png")
	engo.Files.Load("ui/red_circle.png")
	engo.Files.Load("ui/green_tick.png")
	engo.Files.Load("ui/yellow_circle.png")
	engo.Files.Load("ui/green_panel.png")
	engo.Files.Load("ui/about_store.png")

	engo.Files.LoadReaderData("gofont.ttf", bytes.NewReader(gomono.TTF))

	engo.Input.RegisterButton("Exit", engo.KeyEscape)
}

func (s *MainMenuScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)

	common.SetBackground(color.RGBA{R: 0xAD, G: 0xD8, B: 0xE6, A: 0xFF})

	w.AddSystem(&systems.ExitSystem{})

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

	var buttonable *systems.ButtonAble
	var notbuttonable *systems.NotButtonAble
	w.AddSystemInterface(&systems.ButtonSystem{}, buttonable, notbuttonable)

	bgm := struct {
		ecs.BasicEntity
		common.AudioComponent
	}{BasicEntity: ecs.NewBasic()}
	player, _ := common.LoadedPlayer("pleasantcreek.wav")
	bgm.AudioComponent = common.AudioComponent{Player: player}
	w.AddEntity(&bgm)
	player.Play()

	type sprite struct {
		ecs.BasicEntity
		common.RenderComponent
		common.SpaceComponent
	}

	type button struct {
		ecs.BasicEntity
		common.RenderComponent
		common.SpaceComponent
		common.MouseComponent
		systems.ButtonComponent
	}

	bg := sprite{BasicEntity: ecs.NewBasic()}
	bg.Drawable = pixelshader.PixelRegion{}
	bg.SetShader(shaders.BubbleShader)
	bg.SetZIndex(0)
	w.AddEntity(&bg)

	fnt := &common.Font{
		URL:  "ui/BubblegumSans-Regular.ttf",
		FG:   color.RGBA{R: 0xFF, G: 0xFD, B: 0xD0, A: 0xFF},
		Size: 128,
	}
	fnt.CreatePreloaded()
	mainText := sprite{BasicEntity: ecs.NewBasic()}
	mainText.Position = engo.Point{X: 180, Y: 25}
	mainText.Drawable = common.Text{
		Font: fnt,
		Text: "Scrub",
	}
	mainText.SetZIndex(3)
	w.AddEntity(&mainText)

	// main text background
	card := sprite{BasicEntity: ecs.NewBasic()}
	card.Position = engo.Point{X: 125, Y: 25}
	card.Scale = engo.Point{X: 1, Y: 1}
	card.Drawable, _ = common.LoadedSprite("ui/menu_card.png")
	card.SetZIndex(1)
	w.AddEntity(&card)

	// Version Information
	version := sprite{BasicEntity: ecs.NewBasic()}
	version.Position = engo.Point{X: 425, Y: 135}
	version.Drawable = common.Text{
		Font: fnt,
		Text: getVersion(),
	}
	version.Scale = engo.Point{X: 0.2, Y: 0.2}
	version.SetZIndex(3)
	w.AddEntity(&version)

	// Game Start Button
	gameStartButton := button{BasicEntity: ecs.NewBasic()}
	gameStartButton.Position = engo.Point{X: 225, Y: 160}
	gameStartButton.Scale = engo.Point{X: 1, Y: 1}
	gameStartButton.Drawable, _ = common.LoadedSprite("ui/green_button00.png")
	gameStartButton.Width = gameStartButton.Drawable.Width() * gameStartButton.Scale.X
	gameStartButton.Height = gameStartButton.Drawable.Height() * gameStartButton.Scale.Y
	gameStartButton.SetZIndex(2)
	w.AddEntity(&gameStartButton)

	// Game Start Text
	gameStartText := sprite{BasicEntity: ecs.NewBasic()}
	gameStartText.Position = engo.Point{X: 275, Y: 155}
	gameStartText.Drawable = common.Text{
		Font: fnt,
		Text: "Start",
	}
	gameStartText.Scale = engo.Point{X: 0.4, Y: 0.4}
	gameStartText.SetZIndex(3)
	w.AddEntity(&gameStartText)

	// Options Button
	optsButton := button{BasicEntity: ecs.NewBasic()}
	optsButton.Position = engo.Point{X: 225, Y: 285}
	optsButton.Scale = engo.Point{X: 1, Y: 1}
	optsButton.Drawable, _ = common.LoadedSprite("ui/yellow_button00.png")
	optsButton.Width = optsButton.Drawable.Width() * optsButton.Scale.X
	optsButton.Height = optsButton.Drawable.Height() * optsButton.Scale.Y
	optsButton.SetZIndex(2)
	w.AddEntity(&optsButton)

	// Options Text
	optsText := sprite{BasicEntity: ecs.NewBasic()}
	optsText.Position = engo.Point{X: 275, Y: 290}
	optsText.Drawable = common.Text{
		Font: fnt,
		Text: "options",
	}
	optsText.Scale = engo.Point{X: 0.25, Y: 0.25}
	optsText.SetZIndex(3)
	w.AddEntity(&optsText)

	// Back Button
	backButton := button{BasicEntity: ecs.NewBasic()}
	backButton.Position = engo.Point{X: 475, Y: 30}
	backButton.Scale = engo.Point{X: 1, Y: 1}
	backButton.Drawable, _ = common.LoadedSprite("ui/red_boxCross.png")
	backButton.Width = backButton.Drawable.Width() * backButton.Scale.X
	backButton.Height = backButton.Drawable.Height() * backButton.Scale.Y
	backButton.SetZIndex(2)
	backButton.Hidden = true
	w.AddEntity(&backButton)

	// Housekeeping Selection Button
	hs := button{BasicEntity: ecs.NewBasic()}
	hs.Position = engo.Point{X: 150, Y: 65}
	hs.Scale = engo.Point{X: 1, Y: 1}
	hs.Drawable, _ = common.LoadedSprite("ui/red_circle.png")
	hs.Width = hs.Drawable.Width() * hs.Scale.X
	hs.Height = hs.Drawable.Height() * hs.Scale.Y
	hs.SetZIndex(2)
	hs.Hidden = true
	w.AddEntity(&hs)

	// Housekeeping Text
	hsText := sprite{BasicEntity: ecs.NewBasic()}
	hsText.Position = engo.Point{X: 190, Y: 70}
	hsText.Scale = engo.Point{X: 0.15, Y: 0.15}
	hsText.Drawable = common.Text{
		Font: fnt,
		Text: "Housekeeping",
	}
	hsText.SetZIndex(3)
	hsText.Hidden = true
	w.AddEntity(&hsText)

	// SpringCleaning Selection Button
	sc := button{BasicEntity: ecs.NewBasic()}
	sc.Position = engo.Point{X: 340, Y: 65}
	sc.Scale = engo.Point{X: 1, Y: 1}
	sc.Drawable, _ = common.LoadedSprite("ui/red_circle.png")
	sc.Width = sc.Drawable.Width() * sc.Scale.X
	sc.Height = sc.Drawable.Height() * sc.Scale.Y
	sc.SetZIndex(2)
	sc.Hidden = true
	w.AddEntity(&sc)

	// SpringCleaning Text
	scText := sprite{BasicEntity: ecs.NewBasic()}
	scText.Position = engo.Point{X: 380, Y: 70}
	scText.Scale = engo.Point{X: 0.15, Y: 0.15}
	scText.Drawable = common.Text{
		Font: fnt,
		Text: "Spring Cleaning",
	}
	scText.SetZIndex(3)
	scText.Hidden = true
	w.AddEntity(&scText)

	// Level 1 Button
	lvl1OrigPos := engo.Point{X: 235, Y: 130}
	lvl1FinalPos := engo.Point{X: 235, Y: 130}
	lvl1 := button{BasicEntity: ecs.NewBasic()}
	lvl1.Position = lvl1OrigPos
	lvl1.Scale = engo.Point{X: 0.75, Y: 0.75}
	lvl1.Drawable, _ = common.LoadedSprite("ui/yellow_circle.png")
	lvl1.Width = lvl1.Drawable.Width() * lvl1.Scale.X
	lvl1.Height = lvl1.Drawable.Height() * lvl1.Scale.Y
	lvl1.SetZIndex(2)
	lvl1.Hidden = true
	w.AddEntity(&lvl1)
	lvl1Label := sprite{BasicEntity: ecs.NewBasic()}
	lvl1Label.Position = engo.Point{X: 265, Y: 135}
	lvl1Label.Scale = engo.Point{X: 0.15, Y: 0.15}
	lvl1Label.Drawable = common.Text{
		Font: fnt,
		Text: "1",
	}
	lvl1Label.SetZIndex(3)
	lvl1Label.Hidden = true
	w.AddEntity(&lvl1Label)

	// Level 2 Button
	lvl2OrigPos := engo.Point{X: 280, Y: 130}
	lvl2FinalPos := engo.Point{X: 305, Y: 130}
	lvl2 := button{BasicEntity: ecs.NewBasic()}
	lvl2.Position = lvl2OrigPos
	lvl2.Scale = engo.Point{X: 0.75, Y: 0.75}
	lvl2.Drawable, _ = common.LoadedSprite("ui/yellow_circle.png")
	lvl2.Width = lvl2.Drawable.Width() * lvl2.Scale.X
	lvl2.Height = lvl2.Drawable.Height() * lvl2.Scale.Y
	lvl2.SetZIndex(2)
	lvl2.Hidden = true
	w.AddEntity(&lvl2)
	lvl2Label := sprite{BasicEntity: ecs.NewBasic()}
	lvl2Label.Position = engo.Point{X: 310, Y: 135}
	lvl2Label.Scale = engo.Point{X: 0.15, Y: 0.15}
	lvl2Label.Drawable = common.Text{
		Font: fnt,
		Text: "3",
	}
	lvl2Label.SetZIndex(3)
	lvl2Label.Hidden = true
	w.AddEntity(&lvl2Label)

	// Level 3 Button
	lvl3OrigPos := engo.Point{X: 325, Y: 130}
	lvl3FinalPos := engo.Point{X: 390, Y: 130}
	lvl3 := button{BasicEntity: ecs.NewBasic()}
	lvl3.Position = lvl3OrigPos
	lvl3.Scale = engo.Point{X: 0.75, Y: 0.75}
	lvl3.Drawable, _ = common.LoadedSprite("ui/yellow_circle.png")
	lvl3.Width = lvl3.Drawable.Width() * lvl3.Scale.X
	lvl3.Height = lvl3.Drawable.Height() * lvl3.Scale.Y
	lvl3.SetZIndex(2)
	lvl3.Hidden = true
	w.AddEntity(&lvl3)
	lvl3Label := sprite{BasicEntity: ecs.NewBasic()}
	lvl3Label.Position = engo.Point{X: 355, Y: 135}
	lvl3Label.Scale = engo.Point{X: 0.15, Y: 0.15}
	lvl3Label.Drawable = common.Text{
		Font: fnt,
		Text: "5 minutes",
	}
	lvl3Label.SetZIndex(3)
	lvl3Label.Hidden = true
	w.AddEntity(&lvl3Label)

	// Lets Go Button
	lg := button{BasicEntity: ecs.NewBasic()}
	lg.Position = engo.Point{X: 280, Y: 175}
	lg.Scale = engo.Point{X: 1, Y: 1}
	lg.Drawable, _ = common.LoadedSprite("ui/green_panel.png")
	lg.Width = lg.Drawable.Width() * lg.Scale.X
	lg.Height = lg.Drawable.Height() * lg.Scale.Y
	lg.SetZIndex(2)
	lg.Hidden = true
	w.AddEntity(&lg)
	lgLabel := sprite{BasicEntity: ecs.NewBasic()}
	lgLabel.Position = engo.Point{X: 282, Y: 185}
	lgLabel.Scale = engo.Point{X: 0.55, Y: 0.55}
	lgLabel.Drawable = common.Text{
		Font: fnt,
		Text: "Go!",
	}
	lgLabel.SetZIndex(3)
	lgLabel.Hidden = true
	w.AddEntity(&lgLabel)

	// Shop Button
	shop := button{BasicEntity: ecs.NewBasic()}
	shop.Position = engo.Point{X: 225, Y: 225}
	shop.Scale = engo.Point{X: 1, Y: 1}
	shop.Drawable, _ = common.LoadedSprite("ui/yellow_button00.png")
	shop.Width = shop.Drawable.Width() * shop.Scale.X
	shop.Height = shop.Drawable.Height() * shop.Scale.Y
	shop.SetZIndex(2)
	w.AddEntity(&shop)
	shopLabel := sprite{BasicEntity: ecs.NewBasic()}
	shopLabel.Position = engo.Point{X: 280, Y: 225}
	shopLabel.Scale = engo.Point{X: 0.3, Y: 0.3}
	shopLabel.Drawable = common.Text{
		Font: fnt,
		Text: "Shop",
	}
	shopLabel.SetZIndex(3)
	shopLabel.Hidden = false
	w.AddEntity(&shopLabel)

	// Shop Equipment Selection
	eq := button{BasicEntity: ecs.NewBasic()}
	eq.Position = engo.Point{X: 150, Y: 65}
	eq.Scale = engo.Point{X: 1, Y: 1}
	eq.Drawable, _ = common.LoadedSprite("ui/yellow_circle.png")
	eq.Width = eq.Drawable.Width() * eq.Scale.X
	eq.Height = eq.Drawable.Height() * eq.Scale.Y
	eq.SetZIndex(2)
	eq.Hidden = true
	w.AddEntity(&eq)
	eqLabel := sprite{BasicEntity: ecs.NewBasic()}
	eqLabel.Position = offsetPoint(eq.Position, engo.Point{X: 40, Y: 10})
	eqLabel.Scale = engo.Point{X: 0.1, Y: 0.1}
	eqLabel.Drawable = common.Text{
		Font: fnt,
		Text: "Equipment",
	}
	eqLabel.SetZIndex(3)
	eqLabel.Hidden = true
	w.AddEntity(&eqLabel)

	// Shop Consumables Selection
	con := button{BasicEntity: ecs.NewBasic()}
	con.Position = engo.Point{X: 255, Y: 65}
	con.Scale = engo.Point{X: 1, Y: 1}
	con.Drawable, _ = common.LoadedSprite("ui/yellow_circle.png")
	con.Width = con.Drawable.Width() * con.Scale.X
	con.Height = con.Drawable.Height() * con.Scale.Y
	con.SetZIndex(2)
	con.Hidden = true
	w.AddEntity(&con)
	conLabel := sprite{BasicEntity: ecs.NewBasic()}
	conLabel.Position = offsetPoint(con.Position, engo.Point{X: 40, Y: 10})
	conLabel.Scale = engo.Point{X: 0.1, Y: 0.1}
	conLabel.Drawable = common.Text{
		Font: fnt,
		Text: "Consumables",
	}
	conLabel.SetZIndex(3)
	conLabel.Hidden = true
	w.AddEntity(&conLabel)

	// Shop Upgrades Selection
	upg := button{BasicEntity: ecs.NewBasic()}
	upg.Position = engo.Point{X: 375, Y: 65}
	upg.Scale = engo.Point{X: 1, Y: 1}
	upg.Drawable, _ = common.LoadedSprite("ui/yellow_circle.png")
	upg.Width = upg.Drawable.Width() * upg.Scale.X
	upg.Height = upg.Drawable.Height() * upg.Scale.Y
	upg.SetZIndex(2)
	upg.Hidden = true
	w.AddEntity(&upg)
	upgLabel := sprite{BasicEntity: ecs.NewBasic()}
	upgLabel.Position = offsetPoint(upg.Position, engo.Point{X: 40, Y: 10})
	upgLabel.Scale = engo.Point{X: 0.1, Y: 0.1}
	upgLabel.Drawable = common.Text{
		Font: fnt,
		Text: "Upgrades",
	}
	upgLabel.SetZIndex(3)
	upgLabel.Hidden = true
	w.AddEntity(&upgLabel)

	// Shop Info Panel
	ip := sprite{BasicEntity: ecs.NewBasic()}
	ip.Position = engo.Point{X: 525, Y: 28}
	ip.Scale = engo.Point{X: 1, Y: 1}
	ip.Drawable, _ = common.LoadedSprite("ui/about_store.png")
	ip.SetZIndex(1)
	ip.Hidden = true
	w.AddEntity(&ip)

	// Equipment

	// Indicators
	indicatorSelection := 0
	indicatorSelection++
	indicatorSelection--
	indicator := sprite{BasicEntity: ecs.NewBasic()}
	indicator.Position = offsetPoint(hs.Position, engo.Point{X: 9, Y: 9})
	indicator.Scale = engo.Point{X: 1, Y: 1}
	indicator.Drawable, _ = common.LoadedSprite("ui/green_tick.png")
	indicator.SetZIndex(3)
	indicator.Hidden = true
	w.AddEntity(&indicator)
	indicator2Selection := 0
	indicator2 := sprite{BasicEntity: ecs.NewBasic()}
	indicator2.Position = offsetPoint(lvl1.Position, engo.Point{X: 7, Y: 7})
	indicator2.Scale = engo.Point{X: 0.75, Y: 0.75}
	indicator2.Drawable, _ = common.LoadedSprite("ui/green_tick.png")
	indicator2.SetZIndex(3)
	indicator2.Hidden = true
	w.AddEntity(&indicator2)
	indicator3Selection := 0
	indicator3Selection++
	indicator3Selection--
	indicator3 := sprite{BasicEntity: ecs.NewBasic()}
	indicator3.Position = offsetPoint(eq.Position, engo.Point{X: 9, Y: 9})
	indicator3.Drawable, _ = common.LoadedSprite("ui/green_tick.png")
	indicator3.SetZIndex(3)
	indicator3.Hidden = true
	w.AddEntity(&indicator3)

	// OnClick goes on the end so it can hide/unhide all the elements!
	gameStartButton.OnClick = func() {
		gameStartButton.Hidden = true
		gameStartText.Hidden = true
		optsButton.Hidden = true
		optsText.Hidden = true
		version.Hidden = true
		mainText.Hidden = true
		shop.Hidden = true
		shopLabel.Hidden = true
		backButton.Hidden = false
		hs.Hidden = false
		hsText.Hidden = false
		indicator.Hidden = false
		indicator2.Hidden = false
		sc.Hidden = false
		scText.Hidden = false
		lvl1.Hidden = false
		lvl2.Hidden = false
		lvl3.Hidden = false
		lvl1Label.Hidden = false
		lvl2Label.Hidden = false
		lvl3Label.Hidden = false
		lg.Hidden = false
		lgLabel.Hidden = false
	}
	optsButton.OnClick = func() { println("opts!") }
	backButton.OnClick = func() {
		backButton.Hidden = true
		hs.Hidden = true
		hsText.Hidden = true
		indicator.Hidden = true
		indicator2.Hidden = true
		indicator3.Hidden = true
		sc.Hidden = true
		scText.Hidden = true
		lvl1.Hidden = true
		lvl2.Hidden = true
		lvl3.Hidden = true
		lvl1Label.Hidden = true
		lvl2Label.Hidden = true
		lvl3Label.Hidden = true
		lg.Hidden = true
		lgLabel.Hidden = true
		eq.Hidden = true
		eqLabel.Hidden = true
		con.Hidden = true
		conLabel.Hidden = true
		upg.Hidden = true
		upgLabel.Hidden = true
		ip.Hidden = true
		shop.Hidden = false
		shopLabel.Hidden = false
		gameStartButton.Hidden = false
		gameStartText.Hidden = false
		optsButton.Hidden = false
		optsText.Hidden = false
		version.Hidden = false
		mainText.Hidden = false
	}
	hs.OnClick = func() {
		indicator.Position = offsetPoint(hs.Position, engo.Point{X: 9, Y: 9})
		indicatorSelection = 0
		lvl1.Position = lvl1OrigPos
		lvl1Label.Drawable = common.Text{
			Font: fnt,
			Text: "1",
		}
		lvl1Label.Position = offsetPoint(lvl1OrigPos, engo.Point{X: 30, Y: 5})
		lvl2.Position = lvl2OrigPos
		lvl2Label.Drawable = common.Text{
			Font: fnt,
			Text: "3",
		}
		lvl2Label.Position = offsetPoint(lvl2OrigPos, engo.Point{X: 30, Y: 5})
		lvl3.Position = lvl3OrigPos
		lvl3Label.Drawable = common.Text{
			Font: fnt,
			Text: "5 minutes",
		}
		lvl3Label.Position = offsetPoint(lvl3OrigPos, engo.Point{X: 30, Y: 5})
		if indicator2Selection == 2 {
			indicator2.Position = offsetPoint(lvl3.Position, engo.Point{X: 7, Y: 7})
		} else if indicator2Selection == 1 {
			indicator2.Position = offsetPoint(lvl2.Position, engo.Point{X: 7, Y: 7})
		} else {
			indicator2.Position = offsetPoint(lvl1.Position, engo.Point{X: 7, Y: 7})
		}
	}
	sc.OnClick = func() {
		indicator.Position = offsetPoint(sc.Position, engo.Point{X: 9, Y: 9})
		indicatorSelection = 1
		lvl1.Position = lvl1FinalPos
		lvl1Label.Drawable = common.Text{
			Font: fnt,
			Text: "slow",
		}
		lvl1Label.Position = offsetPoint(lvl1FinalPos, engo.Point{X: 30, Y: 5})
		lvl2.Position = lvl2FinalPos
		lvl2Label.Drawable = common.Text{
			Font: fnt,
			Text: "normal",
		}
		lvl2Label.Position = offsetPoint(lvl2FinalPos, engo.Point{X: 30, Y: 5})
		lvl3.Position = lvl3FinalPos
		lvl3Label.Drawable = common.Text{
			Font: fnt,
			Text: "fast",
		}
		lvl3Label.Position = offsetPoint(lvl3FinalPos, engo.Point{X: 30, Y: 5})
		if indicator2Selection == 2 {
			indicator2.Position = offsetPoint(lvl3.Position, engo.Point{X: 7, Y: 7})
		} else if indicator2Selection == 1 {
			indicator2.Position = offsetPoint(lvl2.Position, engo.Point{X: 7, Y: 7})
		} else {
			indicator2.Position = offsetPoint(lvl1.Position, engo.Point{X: 7, Y: 7})
		}
	}
	lvl1.OnClick = func() {
		indicator2.Position = offsetPoint(lvl1.Position, engo.Point{X: 7, Y: 7})
		indicator2Selection = 0
	}
	lvl2.OnClick = func() {
		indicator2.Position = offsetPoint(lvl2.Position, engo.Point{X: 7, Y: 7})
		indicator2Selection = 1
	}
	lvl3.OnClick = func() {
		indicator2.Position = offsetPoint(lvl3.Position, engo.Point{X: 7, Y: 7})
		indicator2Selection = 2
	}
	lg.OnClick = func() { println("lg") }
	shop.OnClick = func() {
		gameStartButton.Hidden = true
		gameStartText.Hidden = true
		optsButton.Hidden = true
		optsText.Hidden = true
		version.Hidden = true
		mainText.Hidden = true
		shop.Hidden = true
		shopLabel.Hidden = true
		backButton.Hidden = false
		eq.Hidden = false
		eqLabel.Hidden = false
		indicator3.Hidden = false
		con.Hidden = false
		conLabel.Hidden = false
		upg.Hidden = false
		upgLabel.Hidden = false
		ip.Hidden = false
		if indicator3Selection == 2 {
			indicator3.Position = offsetPoint(upg.Position, engo.Point{X: 9, Y: 9})
		} else if indicator3Selection == 1 {
			indicator3.Position = offsetPoint(con.Position, engo.Point{X: 9, Y: 9})
		} else {
			indicator3.Position = offsetPoint(eq.Position, engo.Point{X: 9, Y: 9})
		}
	}
	eq.OnClick = func() {
		indicator3.Position = offsetPoint(eq.Position, engo.Point{X: 9, Y: 9})
		indicator3Selection = 0
	}
	con.OnClick = func() {
		indicator3.Position = offsetPoint(con.Position, engo.Point{X: 9, Y: 9})
		indicator3Selection = 1
	}
	upg.OnClick = func() {
		indicator3.Position = offsetPoint(upg.Position, engo.Point{X: 9, Y: 9})
		indicator3Selection = 3
	}
}

func getVersion() string {
	versions := engo.GetApplicationVersion()
	return "v" + strconv.Itoa(versions[0]) + "." + strconv.Itoa(versions[1]) + "." + strconv.Itoa(versions[2])
}

func offsetPoint(p, o engo.Point) engo.Point {
	return *p.Add(o)
}
