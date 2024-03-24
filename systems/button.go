package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type ButtonComponent struct {
	OnClick func()
}

func (c *ButtonComponent) GetButtonComponent() *ButtonComponent { return c }

type ButtonFace interface {
	GetButtonComponent() *ButtonComponent
}

type ButtonAble interface {
	common.BasicFace
	common.RenderFace
	common.SpaceFace
	common.MouseFace
	ButtonFace
}

type NotButtonComponent struct{}

func (n *NotButtonComponent) GetNotButtonComponent() *NotButtonComponent { return n }

type NotButtonFace interface {
	GetNotButtonComponent() *NotButtonComponent
}

type NotButtonAble interface {
	common.BasicFace
	NotButtonFace
}

type buttonEntity struct {
	*ecs.BasicEntity

	*common.RenderComponent
	*common.SpaceComponent
	*common.MouseComponent

	*ButtonComponent
}

type ButtonSystem struct {
	entities []buttonEntity
}

func (s *ButtonSystem) Add(basic *ecs.BasicEntity, render *common.RenderComponent, space *common.SpaceComponent, mouse *common.MouseComponent, button *ButtonComponent) {
	s.entities = append(s.entities, buttonEntity{basic, render, space, mouse, button})
}

func (s *ButtonSystem) AddByInterface(i ecs.Identifier) {
	if o, ok := i.(ButtonAble); ok {
		s.Add(o.GetBasicEntity(), o.GetRenderComponent(), o.GetSpaceComponent(), o.GetMouseComponent(), o.GetButtonComponent())
	}
}

func (s *ButtonSystem) Remove(basic ecs.BasicEntity) {
	var delete = -1
	for index, entity := range s.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *ButtonSystem) Update(dt float32) {
	for _, e := range s.entities {
		if e.Clicked {
			e.OnClick()
		}
	}
}
