package systems

import (
	"fmt"
	components "github.com/cheng81/go-artemis-spaceshipwarriors/components"
	util "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	a "github.com/cheng81/go-artemis/core"
	am "github.com/cheng81/go-artemis/managers"
	as "github.com/cheng81/go-artemis/systems"
	au "github.com/cheng81/go-artemis/util"
	"math"
)

type CollisionHandler func(*a.Entity, *a.Entity)

func ents(b au.ImmutableBag, f bool) au.ImmutableBag {
	if f {
		return b
	} else {
		return au.EmptyBag()
	}
}

func NewCollisionPair(world *a.World, grp1, grp2 string, handler CollisionHandler) *CollisionPair {
	mgr := world.ManagerOfType(am.GroupManagerTypeId).(*am.GroupManager)
	return &CollisionPair{
		groupEntitiesA: grp1, //ents(mgr.Entities(grp1)),
		groupEntitiesB: grp2, //ents(mgr.Entities(grp2)),
		mgr:            mgr,
		handler:        handler,
	}
}

type CollisionPair struct {
	groupEntitiesA, groupEntitiesB string //au.ImmutableBag
	mgr                            *am.GroupManager
	handler                        CollisionHandler
}

func (c *CollisionPair) checkForCollisions() {
	entsA := ents(c.mgr.Entities(c.groupEntitiesA))
	entsB := ents(c.mgr.Entities(c.groupEntitiesB))

	entsA.ForEach(func(_ int, ai interface{}) {
		ent_a := ai.(*a.Entity)
		entsB.ForEach(func(_ int, bi interface{}) {
			ent_b := bi.(*a.Entity)
			if isCollision(ent_a, ent_b) {
				c.handler(ent_a, ent_b)
			}
		})
	})
}

func isCollision(e1 *a.Entity, e2 *a.Entity) bool {
	p1 := components.GetPosition(e1)
	p2 := components.GetPosition(e2)
	b1 := components.GetBounds(e1)
	b2 := components.GetBounds(e2)

	return distance(p1.X, p1.Y, p2.X, p2.Y)-b1.Radius() < b2.Radius()
}

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func NewCollisionSystem() *a.EntitySystem {
	cs := &CollisionSystem{
		BaseProcessor:  as.NewBaseProcessor(),
		collisionPairs: au.NewBag(64),
	}
	es := a.NewEntitySystem(a.AspectFor(components.PositionType, components.BoundsType), CollisionSystemType, cs)
	return es
}

type CollisionSystem struct {
	*as.BaseProcessor
	collisionPairs *au.Bag
}

func (c *CollisionSystem) Initialize() {
	fmt.Println("CollisionSystem.Initialize")
	c.collisionPairs.Add(NewCollisionPair(c.World(), util.Group_PlayerBullets, util.Group_EnemyShips,
		func(bullet *a.Entity, ship *a.Entity) {
			bp := components.GetPosition(bullet)
			util.EntityExplosion(c.World(), bp.X, bp.Y, 0.1).AddToWorld()
			for i := 0; i < 50; i++ {
				util.EntityParticle(c.World(), bp.X, bp.Y).AddToWorld()
			}
			bullet.DeleteFromWorld()

			health := components.GetHealth(ship)
			pos := components.GetPosition(ship)

			health.Health -= 1
			if health.Health < 0 {
				health.Health = 0
				ship.DeleteFromWorld()
				util.EntityExplosion(c.World(), pos.X, pos.Y, 0.5).AddToWorld()
			}
		}))
}

func (c *CollisionSystem) ProcessEntities(_ au.ImmutableBag) {
	c.collisionPairs.ForEach(func(_ int, cp interface{}) {
		cp.(*CollisionPair).checkForCollisions()
	})
}
