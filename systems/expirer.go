package systems

import (
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	a "github.com/cheng81/go-artemis/core"
	as "github.com/cheng81/go-artemis/systems"
)

func NewExpirerSystem() *a.EntitySystem {
	return as.NewDelayedEntitySystem(ExpirerSystemType, a.AspectFor(components.ExpiresType), newExpirer())
}

type expirerProcessor struct{}

func newExpirer() as.DelayedProcessor {
	return expirerProcessor(struct{}{})
}

func (_ expirerProcessor) RemainingDelay(_ *as.DelayedEntityProcessor, e *a.Entity) float64 {
	return components.GetExpires(e).Delay()
}
func (_ expirerProcessor) ProcessDelta(_ *as.DelayedEntityProcessor, e *a.Entity, accumDelta float64) {
	expires := components.GetExpires(e)
	expires.Set(expires.Delay() - accumDelta)
}
func (_ expirerProcessor) ProcessExpired(_ *as.DelayedEntityProcessor, e *a.Entity) {
	e.DeleteFromWorld()
}
