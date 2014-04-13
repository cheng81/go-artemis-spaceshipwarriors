package systems

import (
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	util "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	a "github.com/cheng81/go-artemis/core"
	// am "github.com/cheng81/go-artemis/managers"
	as "github.com/cheng81/go-artemis/systems"
	au "github.com/cheng81/go-artemis/util"
)

const (
	half_w = float64(util.FRAME_WIDTH / 2)
	half_h = float64(util.FRAME_HEIGHT / 2)
)

func NewEntitySpawnerSystem() *a.EntitySystem {
	return a.NewEntitySystem(
		a.AspectEmpty(),
		EntitySpawnerSystemType,
		newEntitySpawner())
}

func newEntitySpawner() (out a.EntitySystemProcessor) {
	timer1 := au.NewTimer(2, true)
	timer2 := au.NewTimer(6, true)
	timer3 := au.NewTimer(12, true)

	sproc := func(es *a.EntitySystem) {

		delta := es.World().Delta()
		timer1.Update(delta)
		timer2.Update(delta)
		timer3.Update(delta)
	}

	vp := as.NewVoidProcessor(sproc)

	timer1.SetCallback(func() {
		util.EntityEnemyShip(vp.World(), "enemy1", components.Layer_ACTORS3, 10, util.Randf(-half_w, half_w), half_h+50., 0., -40, 20).AddToWorld()
	})
	timer2.SetCallback(func() {
		util.EntityEnemyShip(vp.World(), "enemy2", components.Layer_ACTORS2, 20, util.Randf(-half_w, half_w), half_h+100., 0., -30, 40).AddToWorld()
	})
	timer3.SetCallback(func() {
		util.EntityEnemyShip(vp.World(), "enemy3", components.Layer_ACTORS1, 60, util.Randf(-half_w, half_w), half_h+200., 0., -20, 70).AddToWorld()
	})

	return vp
}
