package systems

import (
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	as "github.com/cheng81/go-artemis/systems"
)

func NewMovementSystem() (out *a.EntitySystem) {
	out = as.NewEntitySystem(
		MovementSystemType,
		a.AspectFor(components.PositionType, components.VelocityType),
		FuncProcessor(func(e *a.Entity) {
			p := components.GetPosition(e)
			v := components.GetVelocity(e)

			d := out.World().Delta()
			p.X += v.VecX * d
			p.Y += v.VecY * d
		}))
	return
}
