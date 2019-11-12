package chessgo

import (
	"fmt"
	//"github.com/hyunwoo312/chessGo/lib/chessGo"
)

var (
	// Board Definition
	column = [8]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	row    = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	/**
	'P' -> Pawn
	'B' -> Bishop
	'N' -> kNight
	'R' -> Rook
	'Q' -> Queen
	'K' -> King
	*/
	pieces = [6]rune{'P', 'B', 'N', 'R', 'Q', 'K'}
	// Game Definition
	check     rune
	checkmate bool
	threefold bool
	fiftymove bool
)

// Game object representing game's status
type Game struct {
	board [8][8]string
	// black
	//   a b c d e f g h
	// 8 x x x x x x x x 8
	// 7 x x x x x x x x 7
	// 6 x x x x x x x x 6
	// 5 x x x x x x x x 5
	// 4 x x x x x x x x 4
	// 3 x x x x x x x x 3
	// 2 x x x x x x x x 2
	// 1 x x x x x x x x 1
	//   a b c d e f g h
	// white
	turn   int32
	player rune
}

// x->kill, pawn x -> og col to mov col, O-O is castle short, O-O-O is castle long
func (g *Game) makeMove(m string) {
	nm := validMove(m, g)
	if nm == "check" {
		fmt.Println("You are in check.")
	} else if nm == "invalid" {
		fmt.Println("Invalid move.")
	} else {
		// process move. . .
		if g.player == 'w' && !checkmate {
			g.turn++
		}
	}
}

func validMove(m string, g *Game) string {
	b := g.board
	nm := "check"
	return nm
}

////////////////////
func newGame() *Game {
	g := Game{
		[8][8]string{
			/**
			'P' -> Pawn
			'B' -> Bishop
			'N' -> kNight
			'R' -> Rook
			'Q' -> Queen
			'K' -> King
			*/
			{"bR", "bN", "bB", "bQ", "bK", "bB", "bN", "bR"},
			{"bP", "bP", "bP", "bP", "bP", "bP", "bP", "bP"},
			{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
			{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
			{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
			{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
			{"wP", "wP", "wP", "wP", "wP", "wP", "wP", "wP"},
			{"wR", "wN", "wB", "wQ", "wK", "wB", "wN", "wR"},
		},
		0,   //turn
		'w', //side
	}
	return &g
}

////////////////////
func start() {
	gameover := false
	g := newGame()
	for !(gameover) {
		if checkmate || threefold || fiftymove {
			gameover = true
			// last player at default is the winner unless fiftymove
		}
	}
	fmt.Println("Gamd Over")
}
