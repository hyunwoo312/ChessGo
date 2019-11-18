package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	windowWidth  int32  = 800
	windowHeight int32  = 600
	gamefont     string = "fo/default.ttf"
)

var (
	//fontcolor sdl.Color   = sdl.Color{R: 102, G: 255, B: 214, A: 1}
	eventlist []sdl.Event = make([]sdl.Event, 5)
	err       error
)

func main() {
	// os.Args [0] -> executable file

	// convert png/jpg to bmp format for images
	if len(os.Args) > 2 && os.Args[1] == "--tobmp" {
		// preprocess image files (?) if necessary
		err = tobmp()
		if err != nil {
			panic(err)
		}
		fmt.Println("bitmaps successfully created")
		os.Exit(0)
	}

	fmt.Println("Running ChessAI. . .")
	// Initialize SDL and others. . .
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("SDL Initialization error: ", err)
		os.Exit(1)
	}
	if err = ttf.Init(); err != nil {
		fmt.Println("TTF Initialization error: ", err)
		os.Exit(1)
	}
	defer sdl.Quit()
	defer ttf.Quit()

	// Create Window
	window, err := sdl.CreateWindow(
		"ChessAI",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Window Initialization error: ", err)
		os.Exit(2)
	}
	window.SetResizable(true)
	defer window.Destroy()

	// Create Renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer: ", err)
		os.Exit(2)
	}
	defer renderer.Destroy()

	// Create Board
	user, err := newPlayer(renderer, "white", 480)
	if err != nil {
		fmt.Println("Creating player: ", err)
		os.Exit(2)
	}
	cpu, err := newPlayer(renderer, "black", 480)
	if err != nil {
		fmt.Println("logging computer: ", err)
		os.Exit(2)
	}

	// Create Background
	renderer.SetDrawColor(3, 153, 255, 255)
	img, _ := sdl.LoadBMP("img/cutegophers.bmp")
	defer img.Free()
	bg, _ := renderer.CreateTextureFromSurface(img)
	defer bg.Destroy()

	// Initialize Board
	// board, err := newBoard(renderer)
	// if err != nil {
	// 	fmt.Println("creating board: ", err)
	// 	os.Exit(1)
	// }

	// to := make(chan bool, 0)
	// from := make(chan bool, 0)
	running := true
	var xd int32 = 0
	var yd int32 = 0
	for running {
		sdl.PumpEvents()
		num, err := sdl.PeepEvents(eventlist, sdl.PEEKEVENT, sdl.FIRSTEVENT, sdl.LASTEVENT)
		if err != nil {
			fmt.Printf("Peepevents error: %v\n", err)

		} else {
			for i := 0; i < num; i++ {
				fmt.Printf("Event Peeked Value: %v\n", eventlist[i])
			}
		}

		// Manage event. . .
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) { // x:=
			case *sdl.QuitEvent:
				running = false
				fmt.Printf("Exiting at . . . %v\n", e.GetTimestamp())
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					e.Timestamp, e.Type, e.Which, e.X, e.Y, e.XRel, e.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					e.Timestamp, e.Type, e.Which, e.X, e.Y, e.Button, e.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					e.Timestamp, e.Type, e.Which, e.X, e.Y)
			case *sdl.KeyboardEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					e.Timestamp, e.Type, e.Keysym.Sym, e.Keysym.Mod, e.State, e.Repeat)
				if e.Keysym.Scancode == sdl.SCANCODE_RIGHT {
					xd += 3
				}
				if e.Keysym.Scancode == sdl.SCANCODE_UP {
					yd -= 3
				}
				if e.Keysym.Scancode == sdl.SCANCODE_DOWN {
					yd += 3
				}
				if e.Keysym.Scancode == sdl.SCANCODE_LEFT {
					xd -= 3
				}
			}
		}
		renderer.Clear()

		// Show. . .
		renderer.Copy(bg, nil, &sdl.Rect{X: xd, Y: yd, W: 60, H: 60}) // Chessboard background rendering
		user.draw(renderer)                                           // Show User's pieces
		cpu.draw(renderer)                                            // Show Computer's pieces
		renderer.Present()                                            // Update Screen

		//sdl.Delay(1000 / 30) // ~= 30fps
	}
}
