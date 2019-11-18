package main

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	defaultpattern string = ""
)

// Board object aka Game object representing one's game
type Board struct {
	texture *sdl.Texture
	white   Player
	black   Player
	turn    int32
	// for the sake of math, board will be represented in a 1D array
	status chan string
}

func newBoard(renderer *sdl.Renderer, u, c string, s int32) (Board, error) {
	b := Board{}
	// Initialize board texture
	// Call *Board.Init() to initialize the texture. . .

	// Initialize two players
	user, err := newPlayer(renderer, u, s)
	if err != nil {
		return Board{}, fmt.Errorf("creating new user object: %v", err)
	}
	cpu, err := newPlayer(renderer, c, s)
	if err != nil {
		return Board{}, fmt.Errorf("creating new cpu object: %v", err)
	}
	if u == "white" && c == "black" {
		b.white = user
		b.black = cpu
	} else if u == "black" && c == "white" {
		b.black = user
		b.white = cpu
	} else {
		return Board{}, fmt.Errorf("invalid teams: %v", errors.New("check team colors"))
	}

	// turn default 0
	b.turn = 0

	// blocking channel to take any move
	// and log it
	b.status = make(chan string, 0)
	return b, nil
}

// func (b *Board) Init(renderer *sdl.Renderer, size int32, pattern ...string) (err error) {
// 	if len(pattern) == 0 {
// 		def := defaultpattern
// 	}
// 	return nil
// }
