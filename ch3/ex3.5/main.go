package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

func randUint8() uint8 {
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(256))
}

var RED uint8 = randUint8()
var BLUE uint8 = randUint8()
var GREEN uint8 = randUint8()

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, 2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/width*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{RED, GREEN, BLUE, 255 - contrast*n}
		}
	}
	return color.Black
}
