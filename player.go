package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Player object
// contains 16 *sdl.Texture objects of chess pieces
// color for side,
type Player struct {
	tex *sdl.Texture
	//color  rune
	//pieces *Pieces
}

func newPlayer(renderer *sdl.Renderer) (p Player, err error) {
	image, err := sdl.LoadBMP("img/whiteK.bmp")
	if err != nil {
		return Player{}, fmt.Errorf("loading player sprite: %v", err)
	}
	defer image.Free()

	p.tex, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		return Player{}, fmt.Errorf("creating player sprite: %v", err)
	}

	return p, nil
}

func (p *Player) draw(renderer *sdl.Renderer) {
	renderer.Copy(
		p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 60, H: 60},   // section of the image
		&sdl.Rect{X: 0, Y: 480, W: 60, H: 60}) // location/destination wh -> area/size
}

type Pieces struct {
	//
	// Pawns
	aP *sdl.Texture
	bP *sdl.Texture
	cP *sdl.Texture
	dP *sdl.Texture
	eP *sdl.Texture
	fP *sdl.Texture
	gP *sdl.Texture
	hP *sdl.Texture
	// R, N, B, Q, K
	aR *sdl.Texture
	hR *sdl.Texture
	bN *sdl.Texture
	gN *sdl.Texture
	cB *sdl.Texture
	fB *sdl.Texture
	// Q and K have no tile dictinction
	// since they are unique pieces
	Q *sdl.Texture
	K *sdl.Texture
}
