package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	au "github.com/cheng81/go-artemis/util"
)

func NewParticleProcessorSystem() *a.EntitySystem {
	aspect := a.AspectForOne(components.ParticleType, components.ParticlesType)
	return nil
}

type particleProcessor struct {
	es               *a.EntitySystem
	particlesByLayer map[components.Layer]*au.Bag
	particlesEmitter map[components.Layer]*a.Entity
}

func (_ *particleProcessor) CheckProcessing() bool { return true }
func (_ *particleProcessor) Begin()                {}
func (_ *particleProcessor) End()                  {}
func (_ *particleProcessor) Initialize()           {}

func (p *particleProcessor) SetEntitySystem(es *a.EntitySystem) {
	p.es = es
}

func (p *particleProcessor) Inserted(e *a.Entity) {

}

func (p *particleProcessor) Removed(e *a.Entity) {

}

func (p *particleProcessor) ProcessEntities(_ au.ImmutableBag) {

}
