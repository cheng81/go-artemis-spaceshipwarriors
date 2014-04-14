package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	au "github.com/cheng81/go-artemis/util"
)

func NewParticleRendererSystem(win sf.RenderTarget) *a.EntitySystem {
	pr := newParticleRenderer(win)
	return a.NewEntitySystem(
		a.AspectFor(
			components.PositionType,
			components.ParticleType,
			components.ColorType),
		ParticleRendererSystemType,
		pr)
}

func newParticleRenderer(win sf.RenderTarget) *particleRenderer {
	tex := textureOf("particle")
	texSize := sf.Vector2f{
		float32(tex.GetSize().X),
		float32(tex.GetSize().Y),
	}
	return &particleRenderer{
		EntitySystem: nil,
		tex:          tex,
		texSize:      texSize,
		vertices:     newVertexArray(),
		win:          win,
	}
}

type particleRenderer struct {
	*a.EntitySystem
	tex      *sf.Texture
	texSize  sf.Vector2f
	vertices *sf.VertexArray
	win      sf.RenderTarget
}

func (s *particleRenderer) SetEntitySystem(es *a.EntitySystem) {
	s.EntitySystem = es
}

func (s *particleRenderer) CheckProcessing() bool { return true }
func (s *particleRenderer) Begin()                {}
func (s *particleRenderer) End()                  {}
func (s *particleRenderer) Initialize()           {}

func (s *particleRenderer) Inserted(_ *a.Entity) {}
func (s *particleRenderer) Removed(_ *a.Entity)  {}

func (s *particleRenderer) ProcessEntities(entities au.ImmutableBag) {
	vs := s.vertices
	vs.Clear()
	entities.ForEach(func(_ int, ei interface{}) {
		e := ei.(*a.Entity)
		pos := components.GetPosition(e)
		par := components.GetParticle(e)
		col := components.GetColor(e)
		s.addParticle(par, pos, col, vs)
	})
	rs := sf.DefaultRenderStates()
	rs.Texture = s.tex
	vs.Draw(s.win, rs)
}

func (s *particleRenderer) addParticle(par *components.Particle,
	pos *components.Position,
	col *components.Color,
	vs *sf.VertexArray) {
	texSize := s.texSize
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

func newVertexArray() *sf.VertexArray {
	out, err := sf.NewVertexArray()
	if err != nil {
		fmt.Println("Cannot create vertex array", err)
		panic("Cannot create vertex array")
	}
	return out
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
