package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	a "github.com/cheng81/go-artemis"
	c "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
)

const (
	CollisionSystemType = a.EntitySystemTypeId(iota + 10) //start at 10, 0/1 are component and entity manager
	ColorAnimationSystemType
	EntitySpawnerSystemType
	ExpirerSystemType
	MovementSystemType
	ParallaxStarRepeaterSystemType
	OffScreenRemoverSystemType
	ScaleAnimationSystemType
	PlayerInputSystemType
	SpriteRendererSystemType
	ParticleRendererSystemType
)

type FuncProcessor func(*a.Entity)

func (p FuncProcessor) Process(e *a.Entity) {
	p(e)
}

func sfmlPosition(pos *c.Position) sf.Vector2f {
	return sf.Vector2f{
		float32(pos.X + u.HALF_FRAME_W),
		float32(-pos.Y + u.HALF_FRAME_H),
	}
}
