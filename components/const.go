package components

import (
	a "github.com/cheng81/go-artemis/core"
)

const (
	BoundsType = a.ComponentTypeId(iota)
	ColorAnimationType
	EnemyType
	ExpiresType
	HealthType
	ParallaxStarType
	PlayerType
	PositionType
	ScaleAnimationType
	SpriteType
	VelocityType
	ParticleType
)

type Layer uint

const (
	Layer_DEFAULT = Layer(iota)
	Layer_BACKGROUND
	Layer_ACTORS1
	Layer_ACTORS2
	Layer_ACTORS3
	Layer_PARTICLES
)
