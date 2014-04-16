package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	// "fmt"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	au "github.com/cheng81/go-artemis/util"
	"math"
	// "sort"
)

func NewSpriteProcessorSystem() (out *a.EntitySystem) {
	sr := newSpriteProcessor()
	out = a.NewEntitySystem(
		a.AspectFor(components.PositionType, components.SpriteType, components.ColorType, components.SfRenderableType),
		SpriteProcessorSystemType,
		sr)
	return
}

func newSpriteProcessor() a.EntitySystemProcessor {
	out := &spriteProcessor{
		EntitySystem: nil,
		sprites:      make(map[string]*sf.Sprite),
	}
	return out
}

type spriteProcessor struct {
	*a.EntitySystem
	sprites map[string]*sf.Sprite
}

func (_ *spriteProcessor) CheckProcessing() bool { return true }
func (_ *spriteProcessor) Begin()                {}
func (_ *spriteProcessor) End()                  {}
func (_ *spriteProcessor) Initialize()           {}

func (s *spriteProcessor) SetEntitySystem(es *a.EntitySystem) {
	s.EntitySystem = es
}

func (s *spriteProcessor) Inserted(e *a.Entity) {}
func (s *spriteProcessor) Removed(e *a.Entity)  {}

func (s *spriteProcessor) ProcessEntities(entities au.ImmutableBag) {
	// fmt.Println("spriterRenderer.ProcessEntities")
	// for _, e := range entities {
	// s.process(e)
	// }
	entities.ForEach(func(_ int, ei interface{}) {
		s.process(ei.(*a.Entity))
	})
}

func (s *spriteProcessor) process(e *a.Entity) {
	p := components.GetPosition(e)
	sp := components.GetSprite(e)
	col := components.GetColor(e)
	renderable := components.GetSfRenderable(e)
	sprite := renderable.Drawer.(*sf.Sprite) //s.spriteOf(sp.Name)
	// fmt.Println("spriteProcessor.process - sprite loaded", e, sprite)
	sprite.SetColor(sf.Color{R: b(col.R), G: b(col.G), B: b(col.B), A: b(col.A)})
	sprite.SetScale(sf.Vector2f{float32(sp.ScaleX), float32(sp.ScaleY)})
	sprite.SetPosition(sf.Vector2f{
		float32(p.X + u.HALF_FRAME_W),
		float32(-p.Y + u.HALF_FRAME_H)})

	// sprite.Draw(s.win, sf.DefaultRenderStates())
}

// func (s *spriteProcessor) spriteOf(name string) (out *sf.Sprite) {
// 	out, ok := s.sprites[name]
// 	if !ok {
// 		s.sprites[name] = u.LoadSprite(name)
// 		out = sprite
// 	}
// 	return
// }

func b(f float64) byte {
	return byte(math.Max(0., f) * 255)
}
