package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

// Convert image files to bmp files for OpenGL-Go
// System arguments must be image files
// Function will create a bmp file for the image in the same filepath
func tobmp() (err error) {
	items := os.Args[2:]
	var src image.Image
	for _, f := range items {
		// Load image file. . .
		img, err := os.Open(f)
		if err != nil {
			return fmt.Errorf("file load failed: %v", err)
		}
		defer img.Close()

		// Decode image data. . .
		fname := strings.Split(f, ".")
		format := fname[len(fname)-1]
		var index int
		switch format {
		case "png":
			src, err = png.Decode(img)
			index = 3
		case "jpg":
			src, err = jpeg.Decode(img)
			index = 3
		case "jpeg":
			src, err = jpeg.Decode(img)
			index = 4
		default:
			fmt.Println("unsupported dtype; exiting the program")
			os.Exit(1)
		}
		if err != nil {
			return fmt.Errorf("image decoding failed: %v", err)
		}

		// Create a new Image object
		newimage := image.NewRGBA(src.Bounds())
		draw.Draw(
			newimage,
			newimage.Bounds(),
			&image.Uniform{color.Transparent},
			image.Point{},
			draw.Src)
		// Then paste the source image
		draw.Draw(
			newimage,
			newimage.Bounds(),
			src,
			src.Bounds().Min,
			draw.Over)
		// Create output file. . .
		name := f[:len(f)-index]
		out := name + "bmp"
		target, err := os.Create(out)
		if err != nil {
			return fmt.Errorf("creating file failed: %v", err)
		}
		defer target.Close()

		err = bmp.Encode(target, newimage)
		if err != nil {
			return fmt.Errorf("bmp encoding failed: %v", err)
		}
		// Process completed
	}
	// Success
	return nil
}
