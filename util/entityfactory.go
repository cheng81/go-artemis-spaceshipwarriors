package util

import (
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	a "github.com/cheng81/go-artemis/core"
	am "github.com/cheng81/go-artemis/managers"
	// au "github.com/cheng81/go-artemis/util"
)

func EntityPlayer(w *a.World, x, y float64) (e *a.Entity) {
	e = w.CreateEntity()
	e.AddComponent(components.NewPosition(x, y))

	sprite := components.NewSprite("fighter")
	sprite.R = 93. / 255.
	sprite.G = 255. / 255.
	sprite.B = 129. / 255.
	sprite.InLayer = components.Layer_ACTORS3
	e.AddComponent(sprite)

	e.AddComponent(components.NewVelocityZero())
	e.AddComponent(components.NewBounds(43))
	e.AddComponent(components.NewPlayer())

	w.ManagerOfType(am.GroupManagerTypeId).(*am.GroupManager).Add(e, Group_PlayerShip)
	return
}

func EntityPlayerBullet(w *a.World, x, y float64) (e *a.Entity) {
	e = w.CreateEntity()
	e.AddComponent(components.NewPosition(x, y))

	sprite := components.NewSprite("bullet")
	sprite.InLayer = components.Layer_PARTICLES
	e.AddComponent(sprite)

	e.AddComponent(components.NewVelocity(0, 800))
	e.AddComponent(components.NewBounds(5))
	e.AddComponent(components.NewExpires(5))

	w.ManagerOfType(am.GroupManagerTypeId).(*am.GroupManager).Add(e, Group_PlayerBullets)
	return
}

func EntityEnemyShip(w *a.World, name string, layer components.Layer, health, x, y, velocityX, velocityY, boundsRadius float64) (e *a.Entity) {
	e = w.CreateEntity()

	e.AddComponent(components.NewPosition(x, y))
	e.AddComponent(components.NewVelocity(velocityX, velocityY))
	e.AddComponent(components.NewBounds(boundsRadius))
	e.AddComponent(components.NewHealth(health, health))

	sprite := components.NewSprite(name)
	sprite.InLayer = layer
	sprite.R = 255. / 255.
	sprite.G = 0. / 255.
	sprite.B = 142. / 255.
	e.AddComponent(sprite)

	w.ManagerOfType(am.GroupManagerTypeId).(*am.GroupManager).Add(e, Group_EnemyShips)
	return
}

func EntityExplosion(w *a.World, x, y, scale float64) (e *a.Entity) {
	e = w.CreateEntity()

	e.AddComponent(components.NewPosition(x, y))
	e.AddComponent(components.NewExpires(0.5))
	e.AddComponent(components.NewScaleAnimation(scale/100., scale, -3., false, true))

	sprite := components.NewSprite("explosion")
	sprite.ScaleX = scale
	sprite.ScaleY = scale
	sprite.R = 1
	sprite.G = 216. / 255.
	sprite.B = 0
	sprite.A = 0.5
	sprite.InLayer = components.Layer_PARTICLES
	e.AddComponent(sprite)

	return
}

func EntityStar(w *a.World) (e *a.Entity) {
	e = w.CreateEntity()

	e.AddComponent(components.NewPosition(Randf(-HALF_FRAME_W, HALF_FRAME_W), Randf(-HALF_FRAME_H, HALF_FRAME_H)))
	e.AddComponent(components.NewVelocity(0, Randf(10., 60.)))
	e.AddComponent(components.NewParallaxStar())

	e.AddComponent(components.NewColorAnimation(
		[12]float64{
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0.1, 0.5, Randf(0.2, 0.7)},
		[4]bool{false, false, false, true}, true))

	sprite := components.NewSprite("particle")
	sprite.ScaleX = Randf(0., 0.5)
	sprite.ScaleY = sprite.ScaleX
	sprite.A = Randf(0.1, 0.5)
	sprite.InLayer = components.Layer_BACKGROUND
	e.AddComponent(sprite)

	return
}

func EntityParticle(w *a.World, x, y float64) (e *a.Entity) {
	e = w.CreateEntity()

	e.AddComponent(components.NewPosition(x, y)).
		AddComponent(components.NewVelocity(Randf(-400, 400), Randf(-400, 400))).
		AddComponent(components.NewExpires(1)).
		AddComponent(components.NewColorAnimation(
		[12]float64{
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 1, -1},
		[4]bool{false, false, false, true}, false))

	pScale := Randf(0.3, 0.6)
	e.AddComponent(components.NewParticle(
		pScale, pScale,
		1., 216./255., 0, 0.5,
		components.Layer_PARTICLES))
	// sprite := components.NewSprite("particle")
	// sprite.ScaleX = Randf(0.3, 0.6)
	// sprite.ScaleY = sprite.ScaleX
	// sprite.R = 1
	// sprite.G = 216. / 255.
	// sprite.B = 0
	// sprite.A = 0.5
	// sprite.InLayer = components.Layer_PARTICLES
	// e.AddComponent(sprite)

	return
}
