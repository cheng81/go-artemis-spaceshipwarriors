package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	au "github.com/cheng81/go-artemis/util"
	"math"
	"sort"
)

func NewSpriteRendererSystem(win sf.RenderTarget) (out *a.EntitySystem) {
	sr := newSpriteRenderer(win)
	out = a.NewEntitySystem(
		a.AspectFor(components.PositionType, components.SpriteType, components.ColorType),
		SpriteRendererSystemType,
		sr)
	return
}

func newSpriteRenderer(win sf.RenderTarget) a.EntitySystemProcessor {
	out := &spriteRenderer{
		EntitySystem: nil,
		win:          win,
		// regionsByEntity: au.NewBag(64),
		sortedEntities: entitiesByLayer(make([]*a.Entity, 0)),
		sprites:        make(map[string]*sf.Sprite),
	}
	return out
}

type spriteRenderer struct {
	*a.EntitySystem
	win sf.RenderTarget
	// regionsByEntity *au.Bag
	sortedEntities entitiesByLayer
	sprites        map[string]*sf.Sprite
}

func (_ *spriteRenderer) CheckProcessing() bool { return true }
func (_ *spriteRenderer) Begin()                {}
func (_ *spriteRenderer) End()                  {}
func (_ *spriteRenderer) Initialize()           {}

func (s *spriteRenderer) SetEntitySystem(es *a.EntitySystem) {
	s.EntitySystem = es
}

func (s *spriteRenderer) Inserted(e *a.Entity) {
	// sprite := components.GetSprite(e)
	// fmt.Println("spriteRenderer.Inserted", e.Id(), sprite.Name)
	// s.regionsByEntity.SetAt(e.Id(), sprite.Name)
	s.sortedEntities = append(s.sortedEntities, e)
	sort.Sort(s.sortedEntities)
}
func (s *spriteRenderer) Removed(e *a.Entity) {
	// s.regionsByEntity.SetAt(e.Id(), nil)
	s.sortedEntities = s.sortedEntities.Remove(e)
}

func (s *spriteRenderer) ProcessEntities(entities au.ImmutableBag) {
	// fmt.Println("spriterRenderer.ProcessEntities")
	for _, e := range s.sortedEntities {
		s.process(e)
	}
}

func (s *spriteRenderer) process(e *a.Entity) {
	p := components.GetPosition(e)
	sp := components.GetSprite(e)
	col := components.GetColor(e)
	sprite := s.spriteOf(sp.Name)
	// fmt.Println("spriteRenderer.process - sprite loaded", e, sprite)
	sprite.SetColor(sf.Color{R: b(col.R), G: b(col.G), B: b(col.B), A: b(col.A)})
	sprite.SetScale(sf.Vector2f{float32(sp.ScaleX), float32(sp.ScaleY)})
	sprite.SetPosition(sf.Vector2f{
		float32(p.X + u.HALF_FRAME_W),
		float32(-p.Y + u.HALF_FRAME_H)})

	// sprite.Draw(s.win, sf.DefaultRenderStates())
}

func (s *spriteRenderer) spriteOf(name string) (out *sf.Sprite) {
	out, ok := s.sprites[name]
	if !ok {
		s.sprites[name] = u.LoadSprite(name)
		out = sprite
	}
	return
}

func b(f float64) byte {
	return byte(math.Max(0., f) * 255)
}
