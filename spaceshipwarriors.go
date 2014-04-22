package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	a "github.com/cheng81/go-artemis"
	c "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	s "github.com/cheng81/go-artemis-spaceshipwarriors/systems"
	u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	am "github.com/cheng81/go-artemis/managers"
	"time"
)

func NewSpaceshipWarrior(win *sf.RenderWindow, tpf time.Duration) (out *spaceshipWarriors) {
	preloadTextures()
	out = &spaceshipWarriors{
		win:          win,
		timePerFrame: tpf,
		world:        a.NewWorld(),
		render:       nil,
	}
	out.Init()
	return
}

type spaceshipWarriors struct {
	win          *sf.RenderWindow
	timePerFrame time.Duration

	world *a.World

	render *a.EntitySystem
}

func (sw *spaceshipWarriors) Init() {
	win := sw.win
	w := sw.world
	w.AddManager(am.NewGroupManager())

	w.AddActiveSystem(s.NewMovementSystem())
	w.AddActiveSystem(s.NewPlayerInputSystem(win))
	w.AddActiveSystem(s.NewCollisionSystem())
	w.AddActiveSystem(s.NewExpirerSystem())
	w.AddActiveSystem(s.NewEntitySpawnerSystem())
	w.AddActiveSystem(s.NewParallaxStarRepeaterSystem())
	w.AddActiveSystem(s.NewColorAnimationSystem())
	w.AddActiveSystem(s.NewScaleAnimationSystem())
	w.AddActiveSystem(s.NewOffScreenRemover())
	w.AddActiveSystem(s.NewSpriteProcessorSystem())
	w.AddActiveSystem(s.NewParticleProcessorSystem())

	// sw.spr = w.AddSystem(s.NewSpriteRendererSystem(win), true)
	// sw.par = w.AddSystem(s.NewParticleRendererSystem(win), true)
	sw.render = w.AddSystem(s.NewSfmlRendererSystem(win), true)

	w.Initialize()

	u.EntityPlayer(w, 0, 0).AddToWorld()

	for i := 0; i < 500; i++ {
		u.EntityStar(w).AddToWorld()
	}

	for i := 0; i < int(c.Layer_count); i++ {
		l := c.Layer(i)
		u.EntityParticleEmitter(w, l).AddToWorld()
	}
}

func preloadTextures() {
	u.LoadTexture(u.TEX_PARTICLE)
	u.LoadTexture(u.TEX_BULLET)
	u.LoadTexture(u.TEX_ENEMY1)
	u.LoadTexture(u.TEX_ENEMY2)
	u.LoadTexture(u.TEX_ENEMY3)
	u.LoadTexture(u.TEX_EXPLOSION)
	u.LoadTexture(u.TEX_FIGHTER)
	u.LoadTexture(u.TEX_STAR)
}

func (sw *spaceshipWarriors) Start() {
	now := time.Now()
	lastPrint := now
	deltaTime := time.Duration(0)
	ticker := time.NewTicker(sw.timePerFrame)
	for sw.win.IsOpen() {
		select {
		case <-ticker.C:
			// consume window events
			for ev := sw.win.PollEvent(); ev != nil; ev = sw.win.PollEvent() {
				if ev.Type() == sf.EventTypeClosed {
					sw.win.Close()
				}
			}

			sw.win.Clear(sf.ColorBlack())
			deltaTime = time.Since(now)
			now = time.Now()

			sw.world.SetDelta(deltaTime.Seconds())
			sw.world.Process()

			sw.render.Process()
			// sw.par.Process()
			sw.win.Display()

			if time.Since(lastPrint) > time.Second {
				lastPrint = now
				active := sw.world.EntityManager().ActiveEntitiesCount()
				added := sw.world.EntityManager().TotalAdded()
				created := sw.world.EntityManager().TotalCreated()
				deleted := sw.world.EntityManager().TotalDeleted()
				fmt.Println(
					"\nStats:\n",
					"Active entities: ", active, "\n",
					"Total added: ", added, "\n",
					"Total created:", created, "\n",
					"Total deleted:", deleted, "\n",
					"Discrepancy active~(created-deleted)", (int64(active) - (int64(created) - int64(deleted))))
			}
			// player := sw.world.EntityById(sw.playerId)
			// fmt.Println("spaceshitWarriors.Init - player has input", player.HasSystem(sw.world.SystemOfType(s.PlayerInputSystemType)))

		}
	}
}
