package systems

import (
	sf "bitbucket.org/krepa098/gosfml2"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	a "github.com/cheng81/go-artemis/core"
	as "github.com/cheng81/go-artemis/systems"
)

const (
	hThrusters = 300.
	hMaxSpeed  = 300.
	vThrusters = 200.
	vMaxSpeed  = 200.
	fireRate   = 0.1
)

func NewPlayerInputSystem(win sf.RenderTarget) (out *a.EntitySystem) {
	// fmt.Println("NewPlayerInputSystem called", PlayerInputSystemType)
	out = as.NewEntitySystem(
		PlayerInputSystemType,
		a.AspectFor(
			components.PositionType,
			components.VelocityType,
			components.PlayerType),
		newPlayerInput())
	return
}

func newPlayerInput() *playerInput {
	return &playerInput{fireRate}
}

type playerInput struct {
	timeToFire float64
}

func (pi *playerInput) Process(e *a.Entity) {
	p := components.GetPosition(e)
	v := components.GetVelocity(e)

	// fmt.Println("playerinput::update", e.Id(), p, v)
	d := e.World().Delta()
	if sf.KeyboardIsKeyPressed(sf.KeyUp) {
		v.VecY = u.Clamp(v.VecY+(d*vThrusters), -vMaxSpeed, vMaxSpeed)
	} else if sf.KeyboardIsKeyPressed(sf.KeyDown) {
		v.VecY = u.Clamp(v.VecY-(d*vThrusters), -vMaxSpeed, vMaxSpeed)
	}
	if sf.KeyboardIsKeyPressed(sf.KeyLeft) {
		v.VecX = u.Clamp(v.VecX-(d*hThrusters), -hMaxSpeed, hMaxSpeed)
	} else if sf.KeyboardIsKeyPressed(sf.KeyRight) {
		v.VecX = u.Clamp(v.VecX+(d*hThrusters), -hMaxSpeed, hMaxSpeed)
	}

	if sf.KeyboardIsKeyPressed(sf.KeySpace) {
		if pi.timeToFire <= 0 {
			u.EntityPlayerBullet(e.World(), p.X-27, p.Y+2).AddToWorld()
			u.EntityPlayerBullet(e.World(), p.X+27, p.Y+2).AddToWorld()
			pi.timeToFire = fireRate
		}
	}
	if pi.timeToFire > 0 {
		pi.timeToFire -= d
		if pi.timeToFire < 0 {
			pi.timeToFire = 0
		}
	}

}
