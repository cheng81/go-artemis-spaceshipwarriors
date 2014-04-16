package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	au "github.com/cheng81/go-artemis/util"
)

func NewParticleProcessorSystem() *a.EntitySystem {
	aspect := a.
		AspectForOne(components.ParticleType, components.ParticlesType).
		All(components.LayeredType)
	return a.NewEntitySystem(aspect, ParticleProcessorSystemType, newParticleProcessor())
}

func newParticleProcessor() *particleProcessor {
	tex := textureOf("particle")
	texSize := sf.Vector2f{
		float32(tex.GetSize().X),
		float32(tex.GetSize().Y),
	}
	return &particleProcessor{
		es:               nil,
		particlesByLayer: make(map[components.Layer]*au.Bag),
		particlesEmitter: make(map[components.Layer]*a.Entity),
		tex:              tex,
		texSize:          texSize,
	}
}

type particleProcessor struct {
	es               *a.EntitySystem
	particlesByLayer map[components.Layer]*au.Bag
	particlesEmitter map[components.Layer]*a.Entity
	tex              *sf.Texture
	texSize          sf.Vector2f
}

func (_ *particleProcessor) CheckProcessing() bool { return true }
func (_ *particleProcessor) Begin()                {}
func (_ *particleProcessor) End()                  {}
func (_ *particleProcessor) Initialize()           {}

func (p *particleProcessor) SetEntitySystem(es *a.EntitySystem) {
	p.es = es
}

func (p *particleProcessor) Inserted(e *a.Entity) {
	layer := components.GetLayered(e)
	if e.HasComponent(components.ParticleType) {
		bag, ok := p.particlesByLayer[layer.InLayer()]
		if !ok {
			bag = au.NewBag(256)
			p.particlesByLayer[layer.InLayer()] = bag
		}
		bag.Add(e)
	} else if e.HasComponent(components.ParticlesType) {
		p.particlesEmitter[layer.InLayer()] = e
	}
}

func (p *particleProcessor) Removed(e *a.Entity) {
	layer := components.GetLayered(e)
	if e.HasComponent(components.ParticleType) {
		p.particlesByLayer[layer.InLayer()].RemoveElem(e)
	} else if e.HasComponent(components.ParticlesType) {
		delete(p.particlesEmitter, layer.InLayer())
	}
}

func (p *particleProcessor) ProcessEntities(_ au.ImmutableBag) {
	for layer, ents := range p.particlesByLayer {
		emitter, ok := p.particlesEmitter[layer]
		if ok {
			renderable := components.GetSfRenderable(emitter)
			renderable.States.Texture = p.tex
			vs := renderable.Drawer.(*sf.VertexArray)
			vs.Clear()
			p.process(vs, ents)
		}
	}
}

func (p *particleProcessor) process(vs *sf.VertexArray, parts *au.Bag) {
	parts.ForEach(func(_ int, ei interface{}) {
		e := ei.(*a.Entity)
		pos := components.GetPosition(e)
		col := components.GetColor(e)
		par := components.GetParticle(e)
		p.addParticle(par, pos, col, vs)
	})
}

func (p *particleProcessor) addParticle(par *components.Particle,
	pos *components.Position,
	col *components.Color,
	vs *sf.VertexArray) {
	texSize := p.texSize
	hSize := sf.Vector2f{
		(float32(par.ScaleX) * texSize.X) / 2.,
		(float32(par.ScaleY) * texSize.Y) / 2.,
	}
	color := sf.Color{R: b(col.R), G: b(col.G), B: b(col.B), A: b(col.A)}
	sfPos := sfmlPosition(pos)

	addVertex(vs, sfPos.X-hSize.X, sfPos.Y-hSize.Y, 0, 0, color)
	addVertex(vs, sfPos.X+hSize.X, sfPos.Y-hSize.Y, texSize.X, 0, color)
	addVertex(vs, sfPos.X+hSize.X, sfPos.Y+hSize.Y, texSize.X, texSize.Y, color)
	addVertex(vs, sfPos.X-hSize.X, sfPos.Y+hSize.Y, 0, texSize.Y, color)
}

func addVertex(vs *sf.VertexArray, worldx, worldy, texx, texy float32, color sf.Color) {
	vertex := sf.Vertex{sf.Vector2f{worldx, worldy}, color, sf.Vector2f{texx, texy}}
	vs.Append(vertex)
}

func textureOf(name string) *sf.Texture {
	fname := fmt.Sprint("resources/textures/", name, ".png")
	tex, err := sf.NewTextureFromFile(fname, nil)
	if err != nil {
		fmt.Println("could not load texture", fname, err)
		panic(err)
	}
	return tex
}
