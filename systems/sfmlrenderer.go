package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	// "fmt"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	// u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	au "github.com/cheng81/go-artemis/util"
	// "math"
	"sort"
)

type entitiesByLayer []*a.Entity

func (e entitiesByLayer) Len() int      { return len(e) }
func (e entitiesByLayer) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e entitiesByLayer) Less(i, j int) bool {
	li := components.GetLayered(e[i])
	lj := components.GetLayered(e[j])
	return li.InLayer() < lj.InLayer()
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

func NewSfmlRendererSystem(win sf.RenderTarget) *a.EntitySystem {
	sfrenderer := newSfmlRenderer(win)
	aspect := a.AspectFor(components.SfRenderableType, components.LayeredType)
	return a.NewEntitySystem(aspect, SfmlRendererSystemType, sfrenderer)
}

func newSfmlRenderer(win sf.RenderTarget) a.EntitySystemProcessor {
	return &sfmlRenderer{nil, win, entitiesByLayer(make([]*a.Entity, 0))}
}

type sfmlRenderer struct {
	*a.EntitySystem
	win            sf.RenderTarget
	sortedEntities entitiesByLayer
}

func (_ *sfmlRenderer) CheckProcessing() bool { return true }
func (_ *sfmlRenderer) Begin()                {}
func (_ *sfmlRenderer) End()                  {}
func (_ *sfmlRenderer) Initialize()           {}

func (s *sfmlRenderer) SetEntitySystem(es *a.EntitySystem) {
	s.EntitySystem = es
}

func (s *sfmlRenderer) Inserted(e *a.Entity) {
	s.sortedEntities = append(s.sortedEntities, e)
	sort.Sort(s.sortedEntities)
}
func (s *sfmlRenderer) Removed(e *a.Entity) {
	s.sortedEntities = s.sortedEntities.Remove(e)
}

func (s *sfmlRenderer) ProcessEntities(entities au.ImmutableBag) {
	for _, e := range s.sortedEntities {
		s.process(e)
	}
}

func (s *sfmlRenderer) process(e *a.Entity) {
	renderable := components.GetSfRenderable(e)
	renderable.Draw(s.win, renderable.States)
}
