package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Player object
// contains 16 *sdl.Texture objects of chess pieces
// color for side, incheck for check status
// call Player.Init(side string) to initialize
type Player struct {
	color   string // "white" or "black"
	pieces  [16]Object
	incheck bool
	// sync.Mutex to lock the object to disable concurrent value interference
}

func newPlayer(renderer *sdl.Renderer, s string, dim int32) (p Player, err error) {
	p.color = s

	names := []string{"P", "P", "P", "P", "P", "P", "P", "P", "R", "N", "B", "Q", "K", "B", "N", "R"}
	for i, piece := range names {
		col := 'a'
		var inc int32 = int32(i % 8)
		col += inc
		p.pieces[i], err = newObject(dim, col, s, piece, renderer)
		if err != nil {
			return Player{}, fmt.Errorf("object import error: %v", err)
		}
	}

	p.incheck = false

	// Initialize success
	return p, nil
}

func (p *Player) draw(renderer *sdl.Renderer) {
	for _, item := range p.pieces {
		x := item.coord[0]
		y := item.coord[1]
		renderer.Copy(
			item.texture,
			&sdl.Rect{X: 0, Y: 0, W: 60, H: 60}, // section of the image
			&sdl.Rect{X: x, Y: y, W: 60, H: 60}) // location/destination wh -> area/size
	}
}

// func (p *Player) update(keys []uint8) {
// 	if keys[sdl.SCANCODE_LEFT] == 1 {

// 	} else if keys[sdl.CONTROLLER_BUTTON_LEFTSTICK] {
// 		*sdl.MouseButtonEvent{}
// 	}
// }
