package systems

import (
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	a "github.com/cheng81/go-artemis/core"
	as "github.com/cheng81/go-artemis/systems"
)

func NewOffScreenRemover() (out *a.EntitySystem) {
	out = as.NewIntervalEntitySystem(
		OffScreenRemoverSystemType,
		a.AspectForAll(components.VelocityType,
			components.PositionType,
			components.HealthType,
			components.BoundsType).Exclude(components.PlayerType),
		5,
		FuncProcessor(func(e *a.Entity) {
			p := components.GetPosition(e)
			b := components.GetBounds(e)

			if p.Y < -half_h-b.Radius() {
				e.DeleteFromWorld()
			}
		}))
	return
}
