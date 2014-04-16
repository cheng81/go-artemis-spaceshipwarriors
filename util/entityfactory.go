package util

import (
	a "github.com/cheng81/go-artemis"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	am "github.com/cheng81/go-artemis/managers"
	// au "github.com/cheng81/go-artemis/util"
)

func setSfmlSprite(name string, layer components.Layer, e *a.Entity) {
	e.AddComponent(components.NewLayered(layer)).
		AddComponent(components.NewSprite(name)).
		AddComponent(components.NewSfRenderable(LoadSprite(name)))
}

func EntityPlayer(w *a.World, x, y float64) (e *a.Entity) {
	e = w.CreateEntity()
	e.AddComponent(components.NewPosition(x, y))
	e.AddComponent(components.NewColor(93./255., 255./255., 129./255., 1.))

	setSfmlSprite("fighter", components.Layer_ACTORS3, e)

	// sprite := components.NewSprite("fighter")
	// sprite.InLayer = components.Layer_ACTORS3
	// e.AddComponent(sprite)

	e.AddComponent(components.NewVelocityZero())
	e.AddComponent(components.NewBounds(43))
	e.AddComponent(components.NewPlayer())

	w.ManagerOfType(am.GroupManagerTypeId).(*am.GroupManager).Add(e, Group_PlayerShip)
	return
}

func EntityPlayerBullet(w *a.World, x, y float64) (e *a.Entity) {
	e = w.CreateEntity()
	e.AddComponent(components.NewPosition(x, y))
	e.AddComponent(components.NewColor(1., 1., 1., 1.))

	// sprite := components.NewSprite("bullet")
	// sprite.InLayer = components.Layer_PARTICLES
	// e.AddComponent(sprite)
	setSfmlSprite("bullet", components.Layer_PARTICLES, e)

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
	e.AddComponent(components.NewColor(255./255., 0., 142./255., 1.))

	setSfmlSprite(name, layer, e)
	// sprite := components.NewSprite(name)
	// sprite.InLayer = layer
	// e.AddComponent(sprite)

	w.ManagerOfType(am.GroupManagerTypeId).(*am.GroupManager).Add(e, Group_EnemyShips)
	return
}

func EntityExplosion(w *a.World, x, y, scale float64) (e *a.Entity) {
	e = w.CreateEntity()

	e.AddComponent(components.NewPosition(x, y))
	e.AddComponent(components.NewExpires(0.5))
	e.AddComponent(components.NewScaleAnimation(scale/100., scale, -3., false, true))
	e.AddComponent(components.NewColor(1., 216./255., 0, 0.5))

	setSfmlSprite("explosion", components.Layer_PARTICLES, e)
	sprite := components.GetSprite(e)
	sprite.ScaleX = scale
	sprite.ScaleY = scale

	return
}

func EntityStar(w *a.World) (e *a.Entity) {
	e = w.CreateEntity()

	e.AddComponent(components.NewPosition(Randf(-HALF_FRAME_W, HALF_FRAME_W), Randf(-HALF_FRAME_H, HALF_FRAME_H)))
	e.AddComponent(components.NewVelocity(0, -Randf(10., 60.)))
	e.AddComponent(components.NewParallaxStar())

	e.AddComponent(components.NewColorAnimation(
		[12]float64{
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0.1, 0.5, Randf(0.2, 0.7)},
		[4]bool{false, false, false, true}, true))

	e.AddComponent(components.NewColor(1., 1., 1., Randf(0.1, 0.5)))

	scale := Randf(0., 0.5)
	e.AddComponent(components.NewParticle(scale, scale)).AddComponent(components.NewLayered(components.Layer_BACKGROUND))

	// setSfmlSprite("particle", components.Layer_BACKGROUND, e)
	// sprite := components.GetSprite(e)
	// sprite.ScaleX = Randf(0., 0.5)
	// sprite.ScaleY = sprite.ScaleX
	// e.AddComponent(sprite)

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

	e.AddComponent(components.NewColor(1., 216./255., 0, 0.5))

	pScale := Randf(0.3, 0.6)
	e.AddComponent(components.NewParticle(
		pScale, pScale))

	e.AddComponent(components.NewLayered(components.Layer_PARTICLES))
	return
}

func EntityParticleEmitter(w *a.World, layer components.Layer) (e *a.Entity) {
	e = w.CreateEntity()
	va := NewVertexArray()

	e.AddComponent(components.NewLayered(layer)).
		AddComponent(components.NewParticles()).
		AddComponent(components.NewSfRenderable(va))

	return
}
