package systems

import (
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	as "github.com/cheng81/go-artemis/systems"
)

func NewParallaxStarRepeaterSystem() (out *a.EntitySystem) {
	out = as.NewIntervalEntitySystem(
		ParallaxStarRepeaterSystemType,
		a.AspectFor(components.ParallaxStarType, components.PositionType),
		1,
		FuncProcessor(func(e *a.Entity) {
			p := components.GetPosition(e)
			if p.Y < -half_h {
				p.Y = half_h
			}
		}))
	return
}
