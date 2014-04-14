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

func NewSfmlRendererSystem(win sf.RenderTarget) *a.EntitySystem {
	aspect := a.AspectFor(components.RenderableType)
}

func newSfmlRenderer(win sf.RenderTarget) a.EntitySystemProcessor {

}

type sfmlRenderer struct {
	*a.EntitySystem
	win            sf.RenderTarget
	sortedEntities entitiesByLayer
}
