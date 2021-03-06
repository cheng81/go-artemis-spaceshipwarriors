package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	util "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	au "github.com/cheng81/go-artemis/util"
)

func NewParticleProcessorSystem() *a.EntitySystem {
	aspect := a.
		AspectForOne(components.ParticleType, components.ParticlesType).
		All(components.LayeredType)
	return a.NewEntitySystem(aspect, ParticleProcessorSystemType, newParticleProcessor())
}

func newParticleProcessor() *particleProcessor {
	tex := util.LoadTexture("particle")
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
			for _, v := range vs.Vertices {
				vPool.checkIn(v)
			}
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
	vertex := vPool.checkOut() //sf.Vertex{sf.Vector2f{worldx, worldy}, color, sf.Vector2f{texx, texy}}
	vertex.Position.X = worldx
	vertex.Position.Y = worldy
	vertex.Color = color
	vertex.TexCoords.X = texx
	vertex.TexCoords.Y = texy
	vs.Append(vertex)
}

var vPool = newVertexPool()

func newVertexPool() *vertexPool {
	return &vertexPool{
		vs: au.NewBag(256),
	}
}

type vertexPool struct {
	vs *au.Bag
}

func (v *vertexPool) checkOut() (out sf.Vertex) {
	if v.vs.Size() > 0 {
		return v.vs.RemoveLast().(sf.Vertex)
	}
	return sf.Vertex{sf.Vector2f{}, sf.Color{}, sf.Vector2f{}}
}
func (vp *vertexPool) checkIn(v sf.Vertex) {
	vp.vs.Add(v)
}
