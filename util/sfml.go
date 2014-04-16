package util

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
)

type Centerer interface {
	GetLocalBounds() sf.FloatRect
	SetOrigin(sf.Vector2f)
}

func CenterOrigin(t Centerer) {
	var bounds = t.GetLocalBounds()
	t.SetOrigin(sf.Vector2f{bounds.Width / float32(2.), bounds.Height / float32(2.)})
}

var textures = make(map[string]*sf.Texture)

func LoadSprite(name string) *sf.Sprite {
	tex, ok := textures[name]
	if !ok {
		var err error = nil
		fname := fmt.Sprint("resources/textures/", name, ".png")
		tex, err = sf.NewTextureFromFile(fname, nil)
		if err != nil {
			fmt.Println("could not load texture", fname, err)
			panic(err)
		}
		textures[name] = tex
		fmt.Println("spriteRenderer.spriteOf - sprite loaded", fname)
	}

	out, err := sf.NewSprite(tex)
	if err != nil {
		fmt.Println("could not create sprite", err)
		panic(err)
	}
	CenterOrigin(out)
	return out
}

func NewVertexArray() *sf.VertexArray {
	out, err := sf.NewVertexArray()
	if err != nil {
		fmt.Println("Cannot create vertex array", err)
		panic("Cannot create vertex array")
	}
	return out
}
