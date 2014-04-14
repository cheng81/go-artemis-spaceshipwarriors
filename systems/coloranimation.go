package systems

import (
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	as "github.com/cheng81/go-artemis/systems"
)

func NewColorAnimationSystem() (out *a.EntitySystem) {
	out = as.NewEntitySystem(
		ColorAnimationSystemType,
		a.AspectFor(components.ColorAnimationType, components.ColorType),
		FuncProcessor(func(e *a.Entity) {
			c := components.GetColorAnim(e)
			col := components.GetColor(e)

			if c.AlphaAnimate {
				col.A += c.AlphaSpeed * out.World().Delta()
				if col.A > c.AlphaMax || col.A < c.AlphaMin {
					if c.Repeat {
						c.AlphaSpeed = -c.AlphaSpeed
					} else {
						c.AlphaAnimate = false
					}
				}
			}
		}))
	return
}
