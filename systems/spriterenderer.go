package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	// "container/list"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	// as "github.com/cheng81/go-artemis/systems"
	"fmt"
	au "github.com/cheng81/go-artemis/util"
	"math"
	"sort"
)

func NewSpriteRendererSystem(win sf.RenderTarget) (out *a.EntitySystem) {
	sr := newSpriteRenderer(win)
	out = a.NewEntitySystem(
		a.AspectFor(components.PositionType, components.SpriteType),
		SpriteRendererSystemType,
		sr)
	return
}

type entitiesByLayer []*a.Entity

func (e entitiesByLayer) Len() int      { return len(e) }
func (e entitiesByLayer) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e entitiesByLayer) Less(i, j int) bool {
	si := components.GetSprite(e[i])
	sj := components.GetSprite(e[j])
	return si.InLayer < sj.InLayer
}
func (ents entitiesByLayer) Remove(e *a.Entity) entitiesByLayer {
	index := -1
	for i, el := range ents {
		if el.Id() == e.Id() {
			index = i
			break
		}
	}
	if index < 0 {
		return ents
	}
	copy(ents[index:], ents[index+1:])
	ents[len(ents)-1] = nil
	ents = ents[:len(ents)-1]
	return ents
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
	sprite := s.spriteOf(sp.Name)
	// fmt.Println("spriteRenderer.process - sprite loaded", e, sprite)
	sprite.SetColor(sf.Color{R: b(sp.R), G: b(sp.G), B: b(sp.B), A: b(sp.A)})
	sprite.SetScale(sf.Vector2f{float32(sp.ScaleX), float32(sp.ScaleY)})
	sprite.SetPosition(sf.Vector2f{
		float32(p.X + u.HALF_FRAME_W),
		float32(-p.Y + u.HALF_FRAME_H)})
	sprite.Draw(s.win, sf.DefaultRenderStates())
}

func (s *spriteRenderer) spriteOf(name string) (out *sf.Sprite) {
	out, ok := s.sprites[name]
	if !ok {
		fname := fmt.Sprint("resources/textures/", name, ".png")
		tex, err := sf.NewTextureFromFile(fname, nil)
		if err != nil {
			fmt.Println("could not load texture", fname, err)
			panic(err)
		}
		sprite, err := sf.NewSprite(tex)
		if err != nil {
			fmt.Println("could not create sprite", err)
			panic(err)
		}
		CenterOrigin(sprite)
		fmt.Println("spriteRenderer.spriteOf - sprite loaded", fname, sprite)
		s.sprites[name] = sprite
		out = sprite
	}
	return
}

func b(f float64) byte {
	return byte(math.Max(0., f) * 255)
}

type Centerer interface {
	GetLocalBounds() sf.FloatRect
	SetOrigin(sf.Vector2f)
}

func CenterOrigin(t Centerer) {
	var bounds = t.GetLocalBounds()
	t.SetOrigin(sf.Vector2f{bounds.Width / float32(2.), bounds.Height / float32(2.)})
}
