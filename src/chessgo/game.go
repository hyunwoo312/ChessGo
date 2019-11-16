package chessgo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	//"github.com/hyunwoo312/chessai/lib/webgo"
)

var (
	// Board Definition
	column = [8]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	row    = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	tiles = [64]string
	pieces = [6]rune{'P', 'B', 'N', 'R', 'Q', 'K'}
	// Game Definition
	check     rune
	checkmate bool
	threefold bool
	fiftymove bool
)

// Game object representing game's status
type Game struct {
	Board [8][8]string
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
	Turn  int32
	Side  rune
	Moves chan string
}

func newGame(mode string) (*Game, error) {
	c := make(chan string)
	switch mode {
	// just your basic chess
	// other game mdoes not implemented as of now
	case "normal":
		for i, col := range column {
			for j, rr := range row {
				tiles[j+i*8] = col+string(rr)
			}
		}
		c := make(chan string, 400)
		g := Game{
			[8][8]string{
				/**
				'P' -> Pawn
				'B' -> Bishop
				'N' -> kNight
				'R' -> Rook
				'Q' -> Queen
				'K' -> King
				b -> black && w -> white
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
			c,   //channel of moves
		}
		return &g, nil
	default:
		// Not Accepted Game Mode
		return nil, errors.New("Illegal Operation: Unknown Game Mode")
	}
}

// x->kill, pawn x -> og col to mov col, O-O is castle short, O-O-O is castle long
func (g *Game) makeMove(m string) error {
	// Move Notation
	// m -> side + piece + location > location
	// make move, then change notation to official
	err := validMove(m, g)
	switch err {
	case nil:
		return nil // Invalid move, no changes made
	default:
		// No errors, proceed. . .
		side := m[0]
		piece := m[1]
		from := m[2:4]
		to := m[4:6]
	}
	return nil
}

func validMove(m string, g *Game) error {
	//b := g.board
	// oh no x in y not for python not in golang pepehands
	if g.Side != m[0] || !(m[1] in pieces) || !(m[2:4] in tiles) || !(m[4:6] in tiles) {
		return errors.New("Undefined literals")
	}
	return nil
}

////////////////////
type Move struct {
	GameID int64
	Turn   int32
	Player rune
}
type MoveConfirmation struct {
	Move      Move
	DeltaMove []Move
	Status    Game
	Error     error
}
type MoveResult struct {
	Status Game
}

////////////////////
type API interface {
	CreateGame(ctx context.Context, mode string) error
	FinishGame(ctx context.Context) error
}
type api struct {
	logger *log.Logger
	db     *sql.DB
}

// make new API
func newAPI(logger *log.Logger, db *sql.DB) API {
	return api{logger, db}
}

//
func (api api) CreateGame(ctx context.Context, mode string) error {
	g, error := newGame(mode)
	if error != nil {
		return error
	} else {

	}
}

func (api api) FinishGame(ctx context.Context) error {
	return nil
}

////////////////////
func start() {
	gameover := false
	g := newAPI()
	for !(gameover) {
		if checkmate || threefold || fiftymove {
			gameover = true
			// last player at default is the winner
		}
	}
	// v , ok <- ch ok is false when all channel received and closed
	close(g.moves)
	fmt.Println("Game Over")
}
