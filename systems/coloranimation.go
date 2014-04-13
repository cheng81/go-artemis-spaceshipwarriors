package systems

import (
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	a "github.com/cheng81/go-artemis/core"
	as "github.com/cheng81/go-artemis/systems"
)

func NewColorAnimationSystem() (out *a.EntitySystem) {
	out = as.NewEntitySystem(
		ColorAnimationSystemType,
		a.AspectFor(components.ColorAnimationType).One(components.SpriteType, components.ParticleType),
		FuncProcessor(func(e *a.Entity) {
			c := components.GetColorAnim(e)
			if e.HasComponent(components.SpriteType) {
				sprite := components.GetSprite(e)
				if c.AlphaAnimate {
					sprite.A += c.AlphaSpeed * out.World().Delta()

					if sprite.A > c.AlphaMax || sprite.A < c.AlphaMin {
						if c.Repeat {
							c.AlphaSpeed = -c.AlphaSpeed
						} else {
							c.AlphaAnimate = false
						}
					}
				}
			} else {
				par := components.GetParticle(e)
				if c.AlphaAnimate {
					par.A += c.AlphaSpeed * out.World().Delta()

					if par.A > c.AlphaMax || par.A < c.AlphaMin {
						if c.Repeat {
							c.AlphaSpeed = -c.AlphaSpeed
						} else {
							c.AlphaAnimate = false
						}
					}
				}
			}
		}))
	return
}
