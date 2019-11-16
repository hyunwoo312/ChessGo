package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	windowWidth  int32 = 1280
	windowHeight int32 = 720
)

func main() {
	// os.Args [0] -> executable file

	// convert png/jpg to bmp format for images
	if len(os.Args) > 2 && os.Args[1] == "--tobmp" {
		// preprocess image files (?) if necessary
		err := tobmp()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("bitmaps successfully created")
		os.Exit(0)
	}

	fmt.Println("Running ChessAI. . .")
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("SDL Initialization error: ", err)
		return
	}
	defer sdl.Quit()

	// Create Window
	window, err := sdl.CreateWindow(
		"ChessAI",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Window Initialization error: ", err)
		return
	}
	window.SetResizable(true)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer: ", err)
		return
	}
	// Also needs to close the renderer. . .
	defer renderer.Destroy()
	//bg, err := img.LoadTexture(renderer, "img/21605.jpg")
	user, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("Creating player: ", err)
		return
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		user.draw(renderer)

		renderer.Present()
	}
}
