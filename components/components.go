package components

import (
	a "github.com/cheng81/go-artemis"
)

// getter functions

func GetBounds(e *a.Entity) Bounds { return e.Component(BoundsType).(Bounds) }
func GetColorAnim(e *a.Entity) *ColorAnimation {
	return e.Component(ColorAnimationType).(*ColorAnimation)
}
func GetExpires(e *a.Entity) *Expires   { return e.Component(ExpiresType).(*Expires) }
func GetHealth(e *a.Entity) *Health     { return e.Component(HealthType).(*Health) }
func GetPosition(e *a.Entity) *Position { return e.Component(PositionType).(*Position) }
func GetScaleAnim(e *a.Entity) *ScaleAnimation {
	return e.Component(ScaleAnimationType).(*ScaleAnimation)
}
func GetSprite(e *a.Entity) *Sprite     { return e.Component(SpriteType).(*Sprite) }
func GetVelocity(e *a.Entity) *Velocity { return e.Component(VelocityType).(*Velocity) }
func GetParticle(e *a.Entity) *Particle { return e.Component(ParticleType).(*Particle) }

func NewBounds(r float64) Bounds { return Bounds(r) }

type Bounds float64

func (b Bounds) Radius() float64           { return float64(b) }
func (_ Bounds) TypeId() a.ComponentTypeId { return BoundsType }

func NewColorAnimation(f [12]float64, ctrls [4]bool, repeat bool) *ColorAnimation {
	return &ColorAnimation{
		f[0], f[1], f[2],
		f[3], f[4], f[5],
		f[6], f[7], f[8],
		f[9], f[10], f[11],
		ctrls[0], ctrls[1], ctrls[2], ctrls[3],
		repeat,
	}
}

type ColorAnimation struct {
	RedMin, RedMax, RedSpeed,
	GreenMin, GreenMax, GreenSpeed,
	BlueMin, BlueMax, BlueSpeed,
	AlphaMin, AlphaMax, AlphaSpeed float64

	RedAnimate, GreenAnimate, BlueAnimate, AlphaAnimate, Repeat bool
}

func (_ *ColorAnimation) TypeId() a.ComponentTypeId { return ColorAnimationType }

func NewEnemy() Enemy { return Enemy(struct{}{}) }

type Enemy struct{}

func (_ Enemy) TypeId() a.ComponentTypeId { return EnemyType }

func NewExpires(delay float64) *Expires {
	out := Expires(delay)
	return &out
}

type Expires float64

func (e *Expires) Set(d float64)             { *e = Expires(d) }
func (e *Expires) Delay() float64            { return float64(*e) }
func (_ *Expires) TypeId() a.ComponentTypeId { return ExpiresType }

func NewHealth(h, mh float64) *Health { return &Health{h, mh} }

type Health struct {
	Health, MaxHealth float64
}

func (_ *Health) TypeId() a.ComponentTypeId { return HealthType }

func NewParallaxStar() ParallaxStar { return ParallaxStar(struct{}{}) }

type ParallaxStar struct{}

func (_ ParallaxStar) TypeId() a.ComponentTypeId { return ParallaxStarType }

func NewPlayer() Player { return Player(struct{}{}) }

type Player struct{}

func (_ Player) TypeId() a.ComponentTypeId { return PlayerType }

func NewPosition(x, y float64) *Position {
	return &Position{x, y}
}

type Position struct {
	X, Y float64
}

func (_ *Position) TypeId() a.ComponentTypeId { return PositionType }

func NewScaleAnimation(min, max, speed float64, repeat, active bool) *ScaleAnimation {
	return &ScaleAnimation{min, max, speed, repeat, active}
}

type ScaleAnimation struct {
	Min, Max, Speed float64
	Repeat, Active  bool
}

func (_ *ScaleAnimation) TypeId() a.ComponentTypeId { return ScaleAnimationType }

func NewSprite(name string) *Sprite {
	return &Sprite{name, 1., 1., 0., 1., 1., 1., 1., Layer_DEFAULT}
}

type Sprite struct {
	Name           string
	ScaleX, ScaleY float64
	Rotation       float64
	R, G, B, A     float64
	InLayer        Layer
}

func (_ *Sprite) TypeId() a.ComponentTypeId { return SpriteType }

func NewVelocityZero() *Velocity         { return NewVelocity(0, 0) }
func NewVelocity(x, y float64) *Velocity { return &Velocity{x, y} }

type Velocity struct {
	VecX, VecY float64
}

func (_ *Velocity) TypeId() a.ComponentTypeId { return VelocityType }

func NewParticle(sx, sy, r, g, b, a float64, l Layer) *Particle {
	return &Particle{sx, sy, r, g, b, a, l}
}

type Particle struct {
	ScaleX, ScaleY float64
	R, G, B, A     float64
	InLayer        Layer
}

func (_ *Particle) TypeId() a.ComponentTypeId { return ParticleType }
