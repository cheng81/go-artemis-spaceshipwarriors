package systems

import (
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	a "github.com/cheng81/go-artemis/core"
	as "github.com/cheng81/go-artemis/systems"
)

func NewScaleAnimationSystem() (out *a.EntitySystem) {
	out = as.NewEntitySystem(
		ScaleAnimationSystemType,
		a.AspectFor(components.ScaleAnimationType),
		FuncProcessor(func(e *a.Entity) {
			sa := components.GetScaleAnim(e)
			if sa.Active {
				d := out.World().Delta()
				s := components.GetSprite(e)
				s.ScaleX += sa.Speed * d

				if s.ScaleX >= sa.Max {
					s.ScaleX = sa.Max
					sa.Active = false
				} else if s.ScaleX < sa.Min {
					s.ScaleX = sa.Min
					sa.Active = false
				}

				s.ScaleY = s.ScaleX
			}
		}))
	return
}
