package systems

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/engo/math"
)

type FlashyTextComponent struct {
	Times      []float32
	TextColors []color.NRGBA

	curidx  int
	previdx int
	curTime float32
}

func (c *FlashyTextComponent) GetFlashyTextComponent() *FlashyTextComponent { return c }

type FlashyTextFace interface {
	GetFlashyTextComponent() *FlashyTextComponent
}

type FlashyTextAble interface {
	common.BasicFace
	common.RenderFace
	FlashyTextFace
}

type flashyTextEntity struct {
	*ecs.BasicEntity

	*common.RenderComponent
	*FlashyTextComponent
}

type FlashyTextSystem struct {
	entities []flashyTextEntity
}

func (s *FlashyTextSystem) Add(basic *ecs.BasicEntity, render *common.RenderComponent, flashy *FlashyTextComponent) {
	flashy.previdx = len(flashy.Times) - 1
	s.entities = append(s.entities, flashyTextEntity{basic, render, flashy})
}

func (s *FlashyTextSystem) Remove(basic ecs.BasicEntity) {}

func (s *FlashyTextSystem) AddByInterface(i ecs.Identifier) {
	if o, ok := i.(FlashyTextAble); ok {
		s.Add(o.GetBasicEntity(), o.GetRenderComponent(), o.GetFlashyTextComponent())
	}
}

func (s *FlashyTextSystem) Update(dt float32) {
	for _, ent := range s.entities {
		ent.curTime += dt

		frac := math.Clamp(ent.curTime/ent.Times[ent.curidx], 0.0001, 0.9999)

		prevColor := ent.TextColors[ent.previdx]
		toColor := ent.TextColors[ent.curidx]
		r0, r1 := prevColor.R, toColor.R
		g0, g1 := prevColor.G, toColor.G
		b0, b1 := prevColor.B, toColor.B
		r0 += uint8((float32(r1) - float32(r0)) * frac)
		g0 += uint8((float32(g1) - float32(g0)) * frac)
		b0 += uint8((float32(b1) - float32(b0)) * frac)
		ent.RenderComponent.Color = color.NRGBA{r0, g0, b0, 0xFF}

		if ent.curTime >= ent.Times[ent.curidx] {
			ent.curidx += 1
			ent.curidx %= len(ent.Times)
			ent.previdx += 1
			ent.previdx %= len(ent.Times)
			ent.curTime = 0
		}
	}
}
