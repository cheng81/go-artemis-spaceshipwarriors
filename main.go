package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	u "github.com/cheng81/go-artemis-spaceshipwarriors/util"
	"runtime"
	"time"
)

type FooBat uint

func main() {
	fmt.Println("Starting game")
	w := sf.NewRenderWindow(
		sf.VideoMode{u.FRAME_WIDTH, u.FRAME_HEIGHT, 32},
		"Spaceship Warriors",
		sf.StyleClose,
		sf.DefaultContextSettings())

	sw := NewSpaceshipWarrior(w, time.Duration(time.Second/30))
	sw.Start()
	// b := c.Bounds(50.)
	// d := c.Expires(60.)
	// fmt.Println("Components", c.BoundsType, c.ColorAnimationType, b.TypeId(), d.Delay(), d.TypeId())
}

func init() {
	runtime.LockOSThread()
}
