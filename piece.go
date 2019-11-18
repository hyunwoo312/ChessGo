package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Object type chess piece
// represents the state of the particular piece
// and its relevant information
type Object struct {
	alive   bool
	who     string
	texture *sdl.Texture
	tile    int32 // board tile ranges from 0(bR) -> 63(wR)
	coord   [2]int32
}

func newObject(dim, col int32, side, piece string, renderer *sdl.Renderer) (o Object, err error) {
	o.alive = true
	o.who = string(col) + piece // ex) bR -> black Rook, wK -> white King

	img, err := sdl.LoadBMP("img/" + side + piece + ".bmp") // ex) img/whiteK.bmp
	if err != nil {
		return Object{}, fmt.Errorf("loading icon: %v", err)
	}
	defer img.Free()
	o.texture, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return Object{}, fmt.Errorf("creating texture: %v", err)
	}

	if side == "white" {
		o.tile += 48
		if piece != "P" {
			o.tile += 8
		}
	} else {
		if piece == "P" {
			o.tile += 8
		}
	}
	// 'a' == 97
	xstep := col - 'a'
	o.tile += xstep

	// coord origin -> top left corner x:0 y:0
	x := dim/8*(o.tile%8) + 0
	y := dim/8*(o.tile/8) + 0
	o.coord = [2]int32{x, y}

	// Initialize success
	return o, nil
}

func (o *Object) move(t string) (m string, err error) {
	return "Rb4", nil
}
